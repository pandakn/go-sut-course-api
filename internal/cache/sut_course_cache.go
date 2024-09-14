package cache

import (
	"fmt"
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"
)

type cacheErr string

func (e cacheErr) Error() string {
	return string(e)
}

const (
	ErrCacheMiss cacheErr = "cache miss: key not found"
)

type ISUTCache struct {
	cache *cache.Cache
}

func NewSUTCache() *ISUTCache {
	return &ISUTCache{
		cache: cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (c *ISUTCache) Get(faculty, coursecode, coursename, semester, acadyear, weekdays, timefrom, timeto string, isFilter bool) ([]byte, error) {
	key := generateCacheKey(faculty,coursecode, coursename, semester, acadyear, weekdays, timefrom, timeto, isFilter)

	val, found := c.cache.Get(key)
	if !found {
		return nil, ErrCacheMiss
	}

	return val.([]byte), nil
}

func (c *ISUTCache) Set(faculty, coursecode, coursename, semester, acadyear, weekdays, timefrom, timeto string, isFilter bool, value []byte, expiration time.Duration) error {
	key := generateCacheKey(faculty, coursecode, coursename, semester, acadyear, weekdays, timefrom, timeto, isFilter)

	c.cache.Set(key, value, expiration)
	return nil
}

func generateCacheKey(faculty, coursecode, coursename, semester, acadyear, weekdays, timefrom, timeto string, isFilter bool) string {
	isFilterStr := strconv.FormatBool(isFilter)

  return fmt.Sprintf("%s:%s:%s:%s:%s:%s:%s:%s:%s", faculty,coursecode, coursename, semester, acadyear, isFilterStr, weekdays, timefrom, timeto)
}
