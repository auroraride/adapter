// Copyright (C) adapter. 2024-present.
//
// Created at 2024-09-27, by liasica

package cabdef

// ExchangeStep 换电步骤
type ExchangeStep uint8

const (
	ExchangeStepOpenEmpty ExchangeStep = iota + 1 // 第一步, 开启空电仓
	ExchangeStepPutInto                           // 第二步, 放入旧电池并关闭仓门
	ExchangeStepOpenFull                          // 第三步, 开启满电仓
	ExchangeStepPutOut                            // 第四步, 取出新电池并关闭仓门
)

func (es ExchangeStep) Is(step ExchangeStep) bool {
	return es == step
}

func (es ExchangeStep) EqualInt(n int) bool {
	return es == ExchangeStep(n)
}

func (es ExchangeStep) Int() int {
	return int(es)
}

func (es ExchangeStep) Uint32() uint32 {
	return uint32(es)
}

func (es ExchangeStep) Description() string {
	switch es {
	case ExchangeStepOpenEmpty:
		return "第一步, 开启空电仓"
	case ExchangeStepPutInto:
		return "第二步, 放入旧电池并关闭仓门"
	case ExchangeStepOpenFull:
		return "第三步, 开启满电仓"
	case ExchangeStepPutOut:
		return "第四步, 取出新电池并关闭仓门"
	}
	return "未知"
}
