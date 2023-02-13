// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-17
// Based on adapter by liasica, magicrolan@qq.com.

package mq

import (
    "github.com/auroraride/adapter/log"
    mqtt "github.com/eclipse/paho.mqtt.golang"
    "go.uber.org/zap"
    "time"
)

type Hub struct {
    Server   string
    ClientID string
    Username string
    Password string

    client    mqtt.Client
    listeners map[string]chan []byte
}

func NewHub(server string, id string, username string, password string) *Hub {
    return &Hub{
        Server:   server,
        ClientID: id,
        Username: username,
        Password: password,
    }
}

func (h *Hub) messagePubHandler(client mqtt.Client, msg mqtt.Message) {
    options := client.OptionsReader()
    client.OptionsReader()
    zap.L().WithOptions(zap.WithCaller(false)).Info(
        "MQTT: 收到消息",
        zap.String("topic", msg.Topic()),
        log.Hex(msg.Payload()),
        zap.String("clientid", options.ClientID()),
    )
}

func (h *Hub) connectHandler(client mqtt.Client) {
    options := client.OptionsReader()
    zap.L().WithOptions(zap.WithCaller(false)).Info(
        "MQTT: 已连接",
        zap.String("clientid", options.ClientID()),
    )
}

func (h *Hub) connectLostHandler(client mqtt.Client, err error) {
    options := client.OptionsReader()
    zap.L().WithOptions(zap.WithCaller(false)).Error(
        "MQTT: 已断开连接",
        zap.Error(err),
        zap.String("clientid", options.ClientID()),
    )
}

func (h *Hub) Run() {
    opts := mqtt.NewClientOptions()
    opts.AddBroker(h.Server)
    opts.SetClientID(h.ClientID)
    if h.Username != "" && h.Password != "" {
        opts.SetUsername(h.Username)
        opts.SetPassword(h.Password)
    }
    opts.SetDefaultPublishHandler(h.messagePubHandler)
    opts.OnConnect = h.connectHandler
    opts.OnConnectionLost = h.connectLostHandler
    client := mqtt.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }

    h.client = client
}

// Publish 发送消息
func (h *Hub) Publish(msg *Message) (err error) {
    token := h.client.Publish(msg.Topic, msg.Qos, msg.Retained, msg.Payload)
    token.WaitTimeout(10 * time.Second)
    return token.Error()
}
