// Copyright (C) cabservd. 2025-present.
//
// Created at 2025-02-06, by liasica

package mem

import (
	"context"
	"strconv"
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
func (m *Memcache) AddCabinetBatteryCache(ctx context.Context, serial string, ordinal int, batterySn string) {
	if batterySn == "" {
		return
	}
	m.db.HSet(ctx, m.CabinetBatteryCacheKey(serial), batterySn, ordinal)
}

// RemoveCabinetBatteryCache 移除电柜电池缓存
func (m *Memcache) RemoveCabinetBatteryCache(ctx context.Context, serial string, batterySn string) {
	m.db.HDel(ctx, m.CabinetBatteryCacheKey(serial), batterySn)
}

// ListCabinetBatteryCache 获取电柜电池缓存
func (m *Memcache) ListCabinetBatteryCache(ctx context.Context, serial string) (list map[string]int) {
	list = make(map[string]int)
	val := m.db.HGetAll(ctx, m.CabinetBatteryCacheKey(serial)).Val()
	for k, v := range val {
		list[k], _ = strconv.Atoi(v)
	}
	return
}

// IsMemberCabinetBatteryCache 电池是否在电柜缓存中
func (m *Memcache) IsMemberCabinetBatteryCache(ctx context.Context, serial string, batterySn string) bool {
	return m.db.HExists(ctx, m.CabinetBatteryCacheKey(serial), batterySn).Val()
}

// CountCabinetBatteryCache 电柜电池缓存数量
func (m *Memcache) CountCabinetBatteryCache(ctx context.Context, serial string) int64 {
	return m.db.HLen(ctx, m.CabinetBatteryCacheKey(serial)).Val()
}
