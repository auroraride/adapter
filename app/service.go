// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-23
// Based on adapter by liasica, magicrolan@qq.com.

package app

import (
	"context"
	"net/http"

	"github.com/auroraride/adapter"
)

type Permission bool

const (
	PermissionRequired    Permission = true
	PermissionNotRequired Permission = false
)

type BaseService struct {
	user *adapter.User
	ctx  context.Context
}

func NewService(params ...any) *BaseService {
	nq := PermissionRequired
	s := &BaseService{
		ctx: context.Background(),
	}
	for _, param := range params {
		switch v := param.(type) {
		case *adapter.User:
			s.user = v
		case Permission:
			nq = v
		case context.Context:
			s.ctx = v
		}
	}
	if s.user == nil && nq {
		Panic(http.StatusUnauthorized, adapter.ErrorUserRequired)
	}
	return s
}

func (s *BaseService) GetContext() context.Context {
	return s.ctx
}

func (s *BaseService) GetUser() *adapter.User {
	return s.user
}
