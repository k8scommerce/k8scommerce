package gcache

import (
	"context"
	"fmt"
	"sync"

	"github.com/localrivet/galaxycache"
	"github.com/mailgun/groupcache/v2"
)

var (
	groups  = map[string]*groupcache.Group{}
	getters = map[string]*getter{}
	_cache  Cache
)

func ToKey(params ...interface{}) string {
	return fmt.Sprintf("%v", params)
}

func NewGCache() Cache {
	if _cache == nil {
		_cache = &cache{}
	}
	return _cache
}

type Cache interface {
	NewGroup(groupName string, cacheBytesSize int64, getter groupcache.GetterFunc) *groupcache.Group
	DestroyGroup(groupName string)
	Get(ctx context.Context, groupName, key string) ([]byte, error)
	Delete(ctx context.Context, groupName, key string)
}

type getter struct {
	cacheBytesSize int64
	getterFunc     groupcache.GetterFunc
}

type cache struct {
	mu sync.Mutex
}

func (c *cache) exists(groupName string) bool {
	_, found := groups[groupName]
	return found
}

func (c *cache) NewGroup(groupName string, cacheBytesSize int64, getterFunc groupcache.GetterFunc) *groupcache.Group {
	if !c.exists(groupName) {
		c.mu.Lock()
		defer c.mu.Unlock()
		getters[groupName] = &getter{
			cacheBytesSize: cacheBytesSize,
			getterFunc:     getterFunc,
		}
		groups[groupName] = groupcache.NewGroup(groupName, getters[groupName].cacheBytesSize, getters[groupName].getterFunc)
	}

	return groups[groupName]
}

func (c *cache) DestroyGroup(groupName string) {
	if !c.exists(groupName) {
		c.mu.Lock()
		defer c.mu.Unlock()
		groups[groupName] = nil
		getters[groupName] = nil
	}
}

func (c *cache) Get(ctx context.Context, groupName, key string) ([]byte, error) {
	if c.exists(groupName) {
		codec := &galaxycache.ByteCodec{}
		groups[groupName].Get(ctx, key, nil)
		b, err := codec.MarshalBinary()
		if err != nil {
			return nil, err
		}
		return b, nil
	}
	return nil, fmt.Errorf("key: %s does not exist in galaxy: %s", key, groupName)
}

func (c *cache) Delete(ctx context.Context, groupName, key string) {
	go func() {
		var err error
		if err = c.delete(ctx, groupName, key); err != nil {
			c.NewGroup(groupName, getters[groupName].cacheBytesSize, getters[groupName].getterFunc)
		}
	}()
}

func (c *cache) delete(ctx context.Context, groupName, key string) error {
	if c.exists(key) {
		c.mu.Lock()
		defer c.mu.Unlock()
		return groups[groupName].Remove(ctx, key)
	}
	return nil
}
