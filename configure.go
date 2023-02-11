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

type Environment string

const (
    Production  Environment = "production"
    Development Environment = "development"
)

func (e Environment) String() string {
    return string(e)
}

func (e Environment) UpperString() string {
    return strings.ToUpper(string(e))
}

func (e Environment) IsDevelopment() bool {
    return e == Development
}

type Configurable interface {
    GetApplication() string
    GetEnvironment() Environment
    SetEnvironment(env Environment)
    GetApiAddress() string
    SetKeyPrefix(prefix string)
    GetKeyPrefix() string
    GetCacheKey(key string) string
}

type Configure struct {
    Application string
    keyPrefix   string
    Environment Environment

    Api struct {
        Bind      string
        BodyLimit string
        RateLimit float64
    }

    Redis struct {
        Address  string
        Username string
        Password string
        DB       int `koanf:"db"`
    }
}

func (c *Configure) GetApplication() string {
    return c.Application
}

func (c *Configure) GetEnvironment() Environment {
    return c.Environment
}

func (c *Configure) SetEnvironment(env Environment) {
    c.Environment = env
}

func (c *Configure) GetApiAddress() string {
    return c.Api.Bind
}

func (c *Configure) SetKeyPrefix(prefix string) {
    c.keyPrefix = prefix
}

func (c *Configure) GetKeyPrefix() string {
    return c.keyPrefix
}

func (c *Configure) GetCacheKey(key string) string {
    return c.keyPrefix + key
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
    _, err = os.Stat(cf)
    if defaultConfig != nil && os.IsNotExist(err) {
        err = os.WriteFile(cf, defaultConfig, 0755)
    }

    if err != nil {
        return
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
        cfg.SetKeyPrefix(ApplicationKey(cfg.GetApplication()))
    }

    if cfg.GetEnvironment() == "" {
        cfg.SetEnvironment(Production)
    }

    return
}

func ApplicationKey(application string) string {
    return strings.ToUpper(application) + ":"
}
