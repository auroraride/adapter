// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-17
// Based on adapter by liasica, magicrolan@qq.com.

package mq

import (
    "github.com/auroraride/adapter"
    mqtt "github.com/eclipse/paho.mqtt.golang"
    "time"
)

type Hub struct {
    Server   string
    ClientID string
    Username string
    Password string

    logger    adapter.Logger
    client    mqtt.Client
    listeners map[string]chan []byte
}

func NewHub(server string, id string, username string, password string, logger adapter.Logger) *Hub {
    return &Hub{
        Server:   server,
        ClientID: id,
        Username: username,
        Password: password,
        logger:   logger,
    }
}

func (h *Hub) messagePubHandler(client mqtt.Client, msg mqtt.Message) {
    h.logger.Infof("[MQTT] ↑ Topic: %s, Payload: %x", msg.Topic(), msg.Payload())
}

func (h *Hub) connectHandler(client mqtt.Client) {
    h.logger.Infof("[MQTT] 已连接")
}

func (h *Hub) connectLostHandler(client mqtt.Client, err error) {
    h.logger.Infof("[MQTT] 已断开连接: %v", err)
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
