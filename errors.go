// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-29
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import "errors"

var (
    ErrorData                = errors.New("数据错误")
    ErrorConfig              = errors.New("配置获取失败")
    ErrorLoginExpired        = errors.New("登录超时")
    ErrorForbidden           = errors.New("禁止访问")
    ErrorInvaildCheckSum     = errors.New("数据和校验失败")
    ErrorNotFound            = errors.New("未找到资源")
    ErrorExpired             = errors.New("已过期")
    ErrorBadRequest          = errors.New("请求参数错误")
    ErrorInternalServer      = errors.New("未知错误")
    ErrorUserRequired        = errors.New("需要用户信息")
    ErrorManagerRequired     = errors.New("需要管理员权限")
    ErrorIncompletePacket    = errors.New("incomplete packet") // 数据包不完整
    ErrorIncorrectPacket     = errors.New("消息错误")
    ErrorParamValidateFailed = errors.New("数据校验失败")
    ErrorMaintain            = errors.New("正在唤醒, 请稍后")

    ErrorCabinetBrand          = errors.New("电柜品牌错误")
    ErrorCabinetSerialRequired = errors.New("电柜序号不存在")
    ErrorCabinetBrandRequired  = errors.New("电柜型号不存在")
    ErrorCabinetNotFound       = errors.New("电柜未找到")
    ErrorCabinetOffline        = errors.New("电柜离线")
    ErrorCabinetInitializing   = errors.New("电柜初始化中")
    ErrorCabinetAbnormal       = errors.New("电柜状态异常")
    ErrorCabinetClientNotFound = errors.New("未找到在线电柜")
    ErrorCabinetNoFully        = errors.New("无可换电池")
    ErrorCabinetNoEmpty        = errors.New("无空仓位")
    ErrorCabinetBusy           = errors.New("电柜忙")
    ErrorCabinetControlParam   = errors.New("电柜控制参数错误")
    ErrorCabinetDoorOpened     = errors.New("有开启中的仓门")

    ErrorScanNotExist = errors.New("任务不存在")
    ErrorScanExpired  = errors.New("任务已过期")

    ErrorExchangeCannot       = errors.New("该仓位不满足换电条件")
    ErrorExchangeFailed       = errors.New("换电失败")
    ErrorExchangeBatteryLost  = errors.New("电池未放入")
    ErrorExchangeBatteryExist = errors.New("电池未取走")

    ErrorBinOpened          = errors.New("仓门已是开启状态")
    ErrorBinDisabled        = errors.New("仓位已是禁用状态")
    ErrorBinEnabled         = errors.New("仓位已是启用状态")
    ErrorBinNotFound        = errors.New("仓位未找到")
    ErrorBinOrdinalRequired = errors.New("仓位序号不存在")
    ErrorBinNotEnough       = errors.New("无足够数量的仓位")
    ErrorBinNotUsable       = errors.New("仓位不可用")

    ErrorOperateTimeout = errors.New("操作超时")
    ErrorOperateNoStep  = errors.New("无后续操作")
    ErrorOperateCommand = errors.New("未知的操作指令")

    ErrorBusiness       = errors.New("业务类型错误")
    ErrorBusinessUnable = errors.New("无法办理该业务")

    ErrorBatteryPutin     = errors.New("放入电池编号不匹配")
    ErrorBatteryNotEnough = errors.New("电池数量不足")
    ErrorBatterySN        = errors.New("电池编码错误")
    ErrorBatteryNotFound  = errors.New("未找到当前绑定的电池信息")
)
