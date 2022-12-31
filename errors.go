// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-29
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import "errors"

var (
    ErrorData                      = errors.New("数据错误")
    ErrorNotFound                  = errors.New("未找到资源")
    ErrorBadRequest                = errors.New("请求参数错误")
    ErrorInternalServer            = errors.New("未知错误")
    ErrorUserRequired              = errors.New("需要用户信息")
    ErrorIncompletePacket          = errors.New("incomplete packet") // 数据包不完整
    ErrorParamValidateFailed       = errors.New("数据校验失败")
    ErrorCabinetSerialRequired     = errors.New("电柜序号不存在")
    ErrorCabinetBrandRequired      = errors.New("电柜型号不存在")
    ErrorCabinetBinOrdinalRequired = errors.New("仓位序号不存在")
    ErrorCabinetNotFound           = errors.New("电柜未找到")
    ErrorCabinetBinNotFound        = errors.New("仓位未找到")
    ErrorCabinetOffline            = errors.New("电柜离线")
    ErrorCabinetInitializing       = errors.New("电柜初始化中")
    ErrorCabinetAbnormal           = errors.New("电柜状态异常")
    ErrorCabinetClientNotFound     = errors.New("未找到在线电柜")
    ErrorCabinetNoFully            = errors.New("无可换电池")
    ErrorCabinetNoEmpty            = errors.New("无空仓位")
    ErrorCabinetBusy               = errors.New("电柜忙")
    ErrorCabinetControlParam       = errors.New("电柜控制参数错误")
    ErrorCabinetDoorOpened         = errors.New("有开启中的仓门")
    ErrorCabinetBinNotUsable       = errors.New("仓位不可用")
    ErrorExchangeTaskNotExist      = errors.New("换电任务不存在")
    ErrorExchangeExpired           = errors.New("换电任务已过期")
    ErrorExchangeCannot            = errors.New("该仓位不满足换电条件")
    ErrorExchangeFailed            = errors.New("换电失败")
    ErrorExchangeTimeOut           = errors.New("换电超时")
    ErrorExchangeBatteryLost       = errors.New("电池未放入")
    ErrorExchangeBatteryExist      = errors.New("电池未取走")
    ErrorOperateTimeout            = errors.New("仓位控制超时")
    ErrorBatteryPutin              = errors.New("放入电池编号不匹配")
    ErrorOperate                   = errors.New("未知的操作指令")
    ErrorBinOpened                 = errors.New("仓门已是开启状态")
    ErrorBinDisabled               = errors.New("仓位已是禁用状态")
    ErrorBinEnabled                = errors.New("仓位已是启用状态")
)
