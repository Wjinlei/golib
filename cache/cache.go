package cache

import (
	"sync"
	"time"
)

// 全局map
var (
	mutex sync.Mutex
	obj1  = make(map[string]interface{}) // 缓存对象,里面存放的是缓存的字符串
	obj2  = make(map[string]int64)       // 缓存时间Map,里面存放的是缓存对象的缓存开始时间
	obj3  = make(map[string]int64)       // 缓存过期时间,里面存放的是缓存对象的过期时间,单位是秒
)

// Set 设置缓存, timeoutSecond = 0 或 < 0 都表示立即过期
func Set(key string, value interface{}, timeoutSecond int64) {
	mutex.Lock()
	obj1[key] = value             // 设置缓存对象
	obj2[key] = time.Now().Unix() // 设置缓存开始时间
	obj3[key] = timeoutSecond     // 设置缓存超时时间,单位秒
	mutex.Unlock()
}

// Get 获取缓存
func Get(key string) interface{} {
	if _, ok := obj1[key]; !ok {
		return nil
	}
	if (time.Now().Unix() - obj2[key]) >= obj3[key] {
		delete(obj1, key)
		delete(obj2, key)
		delete(obj3, key)
		return nil
	}
	return obj1[key]
}
