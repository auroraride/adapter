// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-16
// Based on adapter by liasica, magicrolan@qq.com.

package exhook

import (
	"context"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/auroraride/adapter/log"
)

type MessageReceived func(in *MessagePublishRequest) *Message

// Server is used to implement emqx_exhook_v1.s *Server
type Server struct {
	UnimplementedHookProviderServer

	hooks             []Hook
	OnMessageReceived MessageReceived
}

// OnProviderLoaded 定义需要挂载的钩子列表
func (s *Server) OnProviderLoaded(ctx context.Context, in *ProviderLoadedRequest) (*LoadedResponse, error) {
	var hooks []*HookSpec
	for _, spec := range s.hooks {
		hooks = append(hooks, &HookSpec{Name: spec.String()})
	}
	return &LoadedResponse{Hooks: hooks}, nil
}

func (s *Server) OnProviderUnloaded(ctx context.Context, in *ProviderUnloadedRequest) (*EmptySuccess, error) {
	return &EmptySuccess{}, nil
}

func (s *Server) OnClientConnect(ctx context.Context, in *ClientConnectRequest) (*EmptySuccess, error) {
	return &EmptySuccess{}, nil
}

func (s *Server) OnClientConnack(ctx context.Context, in *ClientConnackRequest) (*EmptySuccess, error) {
	return &EmptySuccess{}, nil
}

func (s *Server) OnClientConnected(ctx context.Context, in *ClientConnectedRequest) (*EmptySuccess, error) {
	return &EmptySuccess{}, nil
}

func (s *Server) OnClientDisconnected(ctx context.Context, in *ClientDisconnectedRequest) (*EmptySuccess, error) {
	return &EmptySuccess{}, nil
}

func (s *Server) OnClientAuthenticate(ctx context.Context, in *ClientAuthenticateRequest) (*ValuedResponse, error) {
	reply := &ValuedResponse{}
	reply.Type = ValuedResponse_STOP_AND_RETURN
	reply.Value = &ValuedResponse_BoolResult{BoolResult: true}
	return reply, nil
}

func (s *Server) OnClientAuthorize(ctx context.Context, in *ClientAuthorizeRequest) (*ValuedResponse, error) {
	reply := &ValuedResponse{}
	reply.Type = ValuedResponse_STOP_AND_RETURN
	reply.Value = &ValuedResponse_BoolResult{BoolResult: true}
	return reply, nil
}

func (s *Server) OnClientSubscribe(ctx context.Context, in *ClientSubscribeRequest) (*EmptySuccess, error) {
	return &EmptySuccess{}, nil
}

func (s *Server) OnClientUnsubscribe(ctx context.Context, in *ClientUnsubscribeRequest) (*EmptySuccess, error) {
	return &EmptySuccess{}, nil
}

func (s *Server) OnSessionCreated(ctx context.Context, in *SessionCreatedRequest) (*EmptySuccess, error) {
	return &EmptySuccess{}, nil
}
func (s *Server) OnSessionSubscribed(ctx context.Context, in *SessionSubscribedRequest) (*EmptySuccess, error) {
	return &EmptySuccess{}, nil
}

func (s *Server) OnSessionUnsubscribed(ctx context.Context, in *SessionUnsubscribedRequest) (*EmptySuccess, error) {
	return &EmptySuccess{}, nil
}

func (s *Server) OnSessionResumed(ctx context.Context, in *SessionResumedRequest) (*EmptySuccess, error) {
	return &EmptySuccess{}, nil
}

func (s *Server) OnSessionDiscarded(ctx context.Context, in *SessionDiscardedRequest) (*EmptySuccess, error) {
	return &EmptySuccess{}, nil
}

func (s *Server) OnSessionTakenover(ctx context.Context, in *SessionTakenoverRequest) (*EmptySuccess, error) {
	return &EmptySuccess{}, nil
}

func (s *Server) OnSessionTerminated(ctx context.Context, in *SessionTerminatedRequest) (*EmptySuccess, error) {
	return &EmptySuccess{}, nil
}

func (s *Server) OnMessagePublish(ctx context.Context, in *MessagePublishRequest) (*ValuedResponse, error) {
	zap.L().WithOptions(zap.WithCaller(false)).Info(
		"EXHOOK: 收到消息",
		zap.String("peerhost", in.Message.Headers["peerhost"]),
		zap.String("topic", in.Message.Topic),
		log.Hex(in.Message.Payload),
	)

	var msg *Message
	if s.OnMessageReceived != nil {
		msg = s.OnMessageReceived(in)
	} else {
		msg = in.Message
	}
	return &ValuedResponse{
		Type:  ValuedResponse_STOP_AND_RETURN,
		Value: &ValuedResponse_Message{Message: msg},
	}, nil
}

func (s *Server) OnMessageDelivered(ctx context.Context, in *MessageDeliveredRequest) (*EmptySuccess, error) {
	zap.L().WithOptions(zap.WithCaller(false)).Info(
		"EXHOOK: 发送消息",
		zap.String("clientid", in.Clientinfo.Clientid),
		zap.String("topic", in.Message.Topic),
		log.Hex(in.Message.Payload),
	)
	return &EmptySuccess{}, nil
}

func (s *Server) OnMessageDropped(ctx context.Context, in *MessageDroppedRequest) (*EmptySuccess, error) {
	return &EmptySuccess{}, nil
}

func (s *Server) OnMessageAcked(ctx context.Context, in *MessageAckedRequest) (*EmptySuccess, error) {
	return &EmptySuccess{}, nil
}

func NewServer(hooks ...Hook) *Server {
	if len(hooks) == 0 {
		panic("钩子数量不能为空")
	}
	return &Server{
		hooks: hooks,
	}
}

func (s *Server) Run(address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		zap.L().WithOptions(zap.WithCaller(false)).Fatal(err.Error())
	}
	zap.L().WithOptions(zap.WithCaller(false)).Info("EXHOOK: 启动 -> " + address)

	gs := grpc.NewServer()
	RegisterHookProviderServer(gs, s)
	zap.L().WithOptions(zap.WithCaller(false)).Fatal(gs.Serve(lis).Error())
}
