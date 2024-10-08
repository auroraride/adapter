// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-01
// Based on adapter by liasica, magicrolan@qq.com.

package pqm

import (
	"sync"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/lib/pq"
	"go.uber.org/zap"

	"github.com/auroraride/adapter"
)

type Channelizer interface {
	GetTableName() string
	GetListenerKey() string
}

type Action string

func (a Action) String() string {
	return string(a)
}

const (
	Update Action = "UPDATE"
	Delete Action = "DELETE"
	Insert Action = "INSERT"
)

type Message[T any] struct {
	Table  string `json:"table"`
	Action Action `json:"action"`
	Data   T      `json:"data"`
	Old    T      `json:"old"`
}

func ParseMessage[T any](b []byte) (message *Message[T], err error) {
	message = new(Message[T])
	err = jsoniter.Unmarshal(b, message)
	if err != nil {
		return
	}
	return
}

type Callback[T Channelizer] func(*Message[T])

type Monitor[T Channelizer] struct {
	dsn string

	// 消息回调
	receiver Callback[T]

	// 监听频道
	channel string

	// 监听器
	// 数据格式为: chan *Message[T] -> key
	listeners *sync.Map
}

func NewMonitor[T Channelizer](dsn string, t T, receiver Callback[T]) *Monitor[T] {
	return &Monitor[T]{
		channel:   t.GetTableName(),
		dsn:       dsn,
		receiver:  receiver,
		listeners: &sync.Map{},
	}
}

func (m *Monitor[T]) GetListenerCount() (n int) {
	m.listeners.Range(func(_, _ any) bool {
		n += 1
		return true
	})
	return
}

func (m *Monitor[T]) RemoveListener(ch chan T) {
	m.listeners.Delete(ch)
	close(ch)
}

func (m *Monitor[T]) SetListener(data T, ch chan T) {
	m.listeners.Store(ch, data.GetListenerKey())
}

func (m *Monitor[T]) GetListeners(data T) (chs []chan T) {
	key := data.GetListenerKey()
	m.listeners.Range(func(v, k any) bool {
		if k == key {
			chs = append(chs, v.(chan T))
		}
		return true
	})
	return
}

func (m *Monitor[T]) sendMessage(message *Message[T]) {
	chs := m.GetListeners(message.Data)
	for _, ch := range chs {
		adapter.ChSafeSend(ch, message.Data)
	}
}

func (m *Monitor[T]) Listen() {
	l := pq.NewListener(m.dsn, 10*time.Second, time.Minute, func(ev pq.ListenerEventType, err error) {
		if err != nil {
			zap.L().WithOptions(zap.WithCaller(false)).Error(
				"[MONITOR] "+m.channel+": 监听错误",
				zap.Error(err),
			)
		}
	})

	err := l.Listen(m.channel)
	if err != nil {
		zap.L().WithOptions(zap.WithCaller(false)).Error(
			"[MONITOR] "+m.channel+": 监听失败",
			zap.Error(err),
		)
	}

	zap.L().WithOptions(zap.WithCaller(false)).Info(
		"[MONITOR] " + m.channel + ": 开始监听...",
	)

	after := time.After(90 * time.Second)
	for {
		select {
		case n := <-l.Notify:
			if n.Channel != m.channel {
				continue
			}

			// fmt.Println("[EVENTS] 收到数据库变动 channel [", n.Channel, "] :")
			// var prettyJSON bytes.Buffer
			// _ = json.Indent(&prettyJSON, []byte(n.Extra), "", "  ")
			// fmt.Println(string(prettyJSON.Bytes()))
			// zap.L().Infof("[MONITOR] [%s] 收到数据库变动: \n%s", m.channel, n.Extra)
			// fmt.Printf("[MONITOR] [%s] 收到数据库变动: %s\n", m.channel, n.Extra)

			// TODO 事件通知
			var message *Message[T]
			message, err = ParseMessage[T]([]byte(n.Extra))
			if err != nil {
				zap.L().WithOptions(zap.WithCaller(false)).Error(
					"[MONITOR] "+m.channel+": 消息解析失败",
					zap.Error(err),
					zap.String("extra", n.Extra),
				)
				continue
			}

			// 回调消息
			go m.receiver(message)

			// 检查是否有其他监听器
			go m.sendMessage(message)

		case <-after:
			go func() {
				_ = l.Ping()
			}()
		}
	}
}
