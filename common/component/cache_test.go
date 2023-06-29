package component

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := NewCache()
	key := "a"
	keyB := "keyB"
	cache.SetWithTime(key, "HelloWorld", 3*time.Second)
	cache.Set(keyB, "HelloWorld-long")
	for true {
		time.Sleep(time.Second)
		t.Log(cache.Get(key))
		t.Log(cache.Get(keyB))
	}
}
