// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-18
// Based on adapter by liasica, magicrolan@qq.com.

package exhook

type Hook string

var (
    HookClientConnect       Hook = "client.connect"
    HookClientConnack       Hook = "client.connack"
    HookClientConnected     Hook = "client.connected"
    HookClientDisconnected  Hook = "client.disconnected"
    HookClientAuthenticate  Hook = "client.authenticate"
    HookClientAuthorize     Hook = "client.authorize"
    HookClientSubscribe     Hook = "client.subscribe"
    HookClientUnsubscribe   Hook = "client.unsubscribe"
    HookSessionCreated      Hook = "session.created"
    HookSessionSubscribed   Hook = "session.subscribed"
    HookSessionUnsubscribed Hook = "session.unsubscribed"
    HookSessionResumed      Hook = "session.resumed"
    HookSessionDiscarded    Hook = "session.discarded"
    HookSessionTakenover    Hook = "session.takenover"
    HookSessionTerminated   Hook = "session.terminated"
    HookMessagePublish      Hook = "message.publish"
    HookMessageDelivered    Hook = "message.delivered"
    HookMessageAcked        Hook = "message.acked"
    HookMessageDropped      Hook = "message.dropped"
)

func (h Hook) String() string {
    return string(h)
}
