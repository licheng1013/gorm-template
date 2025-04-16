package tool

import (
	"fmt"
	"sync"
	"time"
)

// Cache 构建一个缓存组件 注意 - 这里只能在单机下使用, 如果多台服务部署, 缓存服务需要切换成redis进行操作
type Cache struct {
	// 锁
	lock sync.Mutex
	// 缓存数据
	cacheMap map[string]interface{}
	// 延迟删除
	delMap map[string]time.Time
}

// NewCache 构建新的实例
func NewCache() *Cache {
	c := &Cache{cacheMap: make(map[string]interface{}), delMap: make(map[string]time.Time)}
	c.startTask()
	return c
}

// Set 缓存数据并设置过期时间
func (c *Cache) Set(key interface{}, value interface{}) {
	c.SetWithTime(key, value, -1)
}

// SetWithTime 缓存数据并设置过期时间
func (c *Cache) SetWithTime(key interface{}, value interface{}, duration time.Duration) {
	defer c.lock.Unlock()
	c.lock.Lock()
	k := fmt.Sprint(key)
	c.cacheMap[k] = value
	// 为负数时,不设置过期时间
	if duration > 0 {
		c.delMap[k] = time.Now().Add(duration)
	}
}

// Get 获取缓存数据
func (c *Cache) Get(key interface{}) interface{} {
	return c.cacheMap[fmt.Sprint(key)]
}

// Del 删除缓存数据
func (c *Cache) Del(key string) {
	defer c.lock.Unlock()
	c.lock.Lock()
	delete(c.cacheMap, key)
	delete(c.delMap, key)
}

// startTask 开启定时任务
func (c *Cache) startTask() {
	go func() {
		for true {
			for k, v := range c.delMap {
				if time.Now().After(v) {
					c.Del(k)
				}
			}
			time.Sleep(time.Second)
		}
	}()
}
