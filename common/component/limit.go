package component

import (
	"sync"
	"time"
)

// 设计一个功能,传入一个id,访问次数,访问时间内
// 在访问时间超过访问次数,回调id给上层处理

func init() {
	// 每隔多少秒清空map
	go func() {
		for true {
			LimitFunc.clearMap()
			time.Sleep(LimitFunc.time)
		}
	}()
}

var LimitFunc = NewLimit()

type Limit struct {
	// 锁
	lock sync.Mutex
	// map
	mapData map[int64]int
	// 时间内
	time time.Duration
	// 次数
	count int
}

// NewLimit 构造方法
func NewLimit() *Limit {
	return &Limit{
		mapData: make(map[int64]int),
		time:    5 * time.Second,
		count:   60, // 60 / 5 = 12 次/s
	}
}

// 清空map加锁
func (t *Limit) clearMap() {
	defer t.lock.Unlock()
	t.lock.Lock()
	t.mapData = make(map[int64]int)
}

// Add 添加
func (t *Limit) Add(key int64, f func()) {
	defer t.lock.Unlock()
	t.lock.Lock()
	t.mapData[key]++
	// 超过次数,回调
	if t.mapData[key] > t.count {
		f()
	}
}
