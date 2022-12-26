// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-26
// Based on adapter by liasica, magicrolan@qq.com.

package tcp

import (
    "github.com/auroraride/adapter/codec"
    "github.com/auroraride/adapter/logger"
)

type Tcp struct {
    codec  codec.Codec
    logger logger.StdLogger
}
