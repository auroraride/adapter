// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-14
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "github.com/spf13/viper"
    "os"
    "path/filepath"
)

func LoadConfigure[T any](cf string, defaultConfig []byte) (cfg *T, err error) {
    // 判定文件是否存在
    dir := filepath.Dir(cf)
    if _, err = os.Stat(dir); os.IsNotExist(err) {
        err = os.MkdirAll(dir, 0755)
        if err != nil {
            return
        }
    }

    // 写入默认配置
    if _, err = os.Stat(cf); os.IsNotExist(err) {
        _ = os.WriteFile(cf, defaultConfig, 0755)
    }

    // 配置文件和环境变量
    viper.SetConfigFile(cf)
    viper.AutomaticEnv()

    // 读取配置
    err = viper.ReadInConfig()
    if err != nil {
        return
    }

    // 解析配置
    cfg = new(T)
    err = viper.Unmarshal(cfg)
    return
}
