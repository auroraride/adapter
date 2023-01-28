// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-14
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "github.com/knadh/koanf"
    "github.com/knadh/koanf/parsers/yaml"
    "github.com/knadh/koanf/providers/file"
    "github.com/mitchellh/mapstructure"
    "os"
    "path/filepath"
    "strings"
)

type Configurable interface {
    GetApplication() string
    GetApiAddress() string
    SetKeySuffix(suffix string)
    GetCacheKey(key string) string
}

type Configure struct {
    Application string
    KeySuffix   string

    Api struct {
        Bind string
    }

    Redis struct {
        Logkey  string
        Address string
    }
}

func (c *Configure) GetApplication() string {
    return c.Application
}

func (c *Configure) GetApiAddress() string {
    return c.Api.Bind
}

func (c *Configure) SetKeySuffix(suffix string) {
    c.KeySuffix = suffix
}

func (c *Configure) GetCacheKey(key string) string {
    return strings.ToUpper(key) + ":" + c.KeySuffix
}

func LoadConfigure[T Configurable](cfg T, cf string, defaultConfig []byte) (err error) {
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

    k := koanf.New(".")
    f := file.Provider(cf)
    p := yaml.Parser()

    err = k.Load(f, p)
    if err != nil {
        return
    }

    // 解析配置
    err = k.UnmarshalWithConf("", cfg, koanf.UnmarshalConf{DecoderConfig: &mapstructure.DecoderConfig{
        DecodeHook: mapstructure.ComposeDecodeHookFunc(
            mapstructure.StringToTimeDurationHookFunc(),
            mapstructure.StringToSliceHookFunc(","),
            mapstructure.TextUnmarshallerHookFunc()),
        Metadata:         nil,
        Result:           cfg,
        WeaklyTypedInput: true,
        Squash:           true,
    }})

    if err == nil {
        cfg.SetKeySuffix(ApplicationKey(cfg.GetApplication(), cfg.GetApiAddress()))
    }

    return
}

func ApplicationKey(application, addr string) string {
    index := strings.Index(addr, ":")
    return strings.ToUpper(application) + "_" + addr[index+1:]
}
