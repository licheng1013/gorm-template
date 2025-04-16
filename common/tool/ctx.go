package tool

import (
	"bytes"
	"context"
	"runtime"
	"strconv"
	"sync"
)

var LocalData = NewLocalContext()

func (*LocalContext) GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

// SaveCtx 保存上下文
func (l *LocalContext) SaveCtx(data interface{}) {
	defer l.mu.Unlock()
	l.mu.Lock()
	background := context.Background()
	background = context.WithValue(background, l.key, data)
	l.Ctx[l.GetGID()] = background
}

// GetCtx 获取上下文
func (l *LocalContext) GetCtx() interface{} {
	defer l.mu.Unlock()
	l.mu.Lock()
	value := l.Ctx[l.GetGID()].Value(l.key)
	if value == nil {
		panic("请检查是否是登入后的路由")
	}
	return value
}

// ClearCtx 清理上下文
func (l *LocalContext) ClearCtx() {
	// 加锁
	defer l.mu.Unlock()
	l.mu.Lock()
	delete(l.Ctx, l.GetGID())
}

// LocalContext 本地上下文
type LocalContext struct {
	// 上下文
	Ctx map[uint64]context.Context
	// 上下文key
	key string
	// 锁
	mu sync.Mutex
}

// NewLocalContext 创建本地上下文，不得在子协程中调用，否则会无法获取正确的协程id
func NewLocalContext() *LocalContext {
	return &LocalContext{key: "userIdKey", Ctx: make(map[uint64]context.Context), mu: sync.Mutex{}}
}
