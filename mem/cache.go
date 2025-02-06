// Copyright (C) cabservd. 2025-present.
//
// Created at 2025-02-06, by liasica

package mem

import (
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	cacheInstance *Memcache
	cacheOnece    sync.Once
)

type Memcache struct {
	db *redis.Client

	// 电池缓存前缀
	// CACHE:CABINET:BATTERY:[电柜编码 Serial]
	cabinetBatteryCacheKeyPrefix string
}

func CacheSetup(db *redis.Client) {
	cacheOnece.Do(func() {
		cacheInstance = &Memcache{
			db: db,

			cabinetBatteryCacheKeyPrefix: "CACHE:CABINET:BATTERY:",
		}
	})
}

// Cache 获取缓存实例
func Cache() *Memcache {
	return cacheInstance
}

// CabinetBatteryCacheKey 电柜电池缓存键
func (m *Memcache) CabinetBatteryCacheKey(serial string) string {
	return m.cabinetBatteryCacheKeyPrefix + serial
}

// ClearCabinetBatteryCache 清除电柜电池缓存
func (m *Memcache) ClearCabinetBatteryCache(ctx context.Context, serial string) {
	m.db.Del(ctx, m.CabinetBatteryCacheKey(serial))
}

// AddCabinetBatteryCache 添加电柜电池缓存
func (m *Memcache) AddCabinetBatteryCache(ctx context.Context, serial string, batterySn string) {
	if batterySn == "" {
		return
	}
	m.db.SAdd(ctx, m.CabinetBatteryCacheKey(serial), batterySn)
}

// RemoveCabinetBatteryCache 移除电柜电池缓存
func (m *Memcache) RemoveCabinetBatteryCache(ctx context.Context, serial string, batterySn string) {
	m.db.SRem(ctx, m.CabinetBatteryCacheKey(serial), batterySn)
}

// ListCabinetBatteryCache 获取电柜电池缓存
func (m *Memcache) ListCabinetBatteryCache(ctx context.Context, serial string) []string {
	return m.db.SMembers(ctx, m.CabinetBatteryCacheKey(serial)).Val()
}

// IsMemberCabinetBatteryCache 电池是否在电柜缓存中
func (m *Memcache) IsMemberCabinetBatteryCache(ctx context.Context, serial string, batterySn string) bool {
	return m.db.SIsMember(ctx, m.CabinetBatteryCacheKey(serial), batterySn).Val()
}
