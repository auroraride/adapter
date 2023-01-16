// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-16
// Based on adapter by liasica, magicrolan@qq.com.

package exhook

import (
    "context"
)

// Server is used to implement emqx_exhook_v1.s *Server
type Server struct {
    UnimplementedHookProviderServer
}

func (s *Server) OnProviderLoaded(ctx context.Context, in *ProviderLoadedRequest) (*LoadedResponse, error) {
    hooks := []*HookSpec{
        {Name: "client.connect"},
        {Name: "client.connack"},
        {Name: "client.connected"},
        {Name: "client.disconnected"},
        {Name: "client.authenticate"},
        {Name: "client.authorize"},
        {Name: "client.subscribe"},
        {Name: "client.unsubscribe"},
        {Name: "session.created"},
        {Name: "session.subscribed"},
        {Name: "session.unsubscribed"},
        {Name: "session.resumed"},
        {Name: "session.discarded"},
        {Name: "session.takenover"},
        {Name: "session.terminated"},
        {Name: "message.publish"},
        {Name: "message.delivered"},
        {Name: "message.acked"},
        {Name: "message.dropped"},
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
    in.Message.Payload = []byte("hardcode payload by exhook-svr-go :)")
    reply := &ValuedResponse{}
    reply.Type = ValuedResponse_STOP_AND_RETURN
    reply.Value = &ValuedResponse_Message{Message: in.Message}
    return reply, nil
}

func (s *Server) OnMessageDelivered(ctx context.Context, in *MessageDeliveredRequest) (*EmptySuccess, error) {
    return &EmptySuccess{}, nil
}

func (s *Server) OnMessageDropped(ctx context.Context, in *MessageDroppedRequest) (*EmptySuccess, error) {
    return &EmptySuccess{}, nil
}

func (s *Server) OnMessageAcked(ctx context.Context, in *MessageAckedRequest) (*EmptySuccess, error) {
    return &EmptySuccess{}, nil
}
