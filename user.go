// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import "database/sql/driver"

const (
    HeaderUserID   = "X-User-ID"
    HeaderUserType = "X-User-Type"
)

type UserType string

const (
    UserTypeUnknown  UserType = "unknown"  // 未知
    UserTypeCabinet  UserType = "cabinet"  // 电柜
    UserTypeManager  UserType = "manager"  // 后台
    UserTypeEmployee UserType = "employee" // 员工
    UserTypeRider    UserType = "rider"    // 骑手
)

type User struct {
    Type UserType `json:"type" validate:"required"` // 用户类别
    ID   string   `json:"id" validate:"required"`   // 用户ID(通常是电话), 电柜的时候使用电柜ID
}

func (t UserType) String() string {
    return string(t)
}

func (t *UserType) Scan(src interface{}) error {
    switch v := src.(type) {
    case nil:
        return nil
    case string:
        *t = UserType(v)
    }
    return nil
}

func (t UserType) Value() (driver.Value, error) {
    return t, nil
}

func (u *User) String() string {
    return u.Type.String() + " - " + u.ID
}
