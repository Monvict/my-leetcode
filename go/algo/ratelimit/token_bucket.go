package main

import (
	"fmt"
	"sync"
	"time"
)

// 令牌桶的原理是：
// 在一个固定容量的桶中，按固定的速率加入token
// 请求时，判断剩余的tokens是否大于0；是，则放行，否则，限流
type TokenBucket struct {
	capacity     float64   // 容量
	fillRate     float64   // 放令牌的速率（个/秒）
	tokens       float64   // 剩余token数
	lastFillTime time.Time // 最后一次填充事件
	mu           sync.Mutex
}

func NewTokenBucket(cap float64, fillRate float64) *TokenBucket {
	return &TokenBucket{
		capacity:     cap,
		fillRate:     fillRate,
		tokens:       cap,
		lastFillTime: time.Now(),
		mu:           sync.Mutex{},
	}
}

// 加入令牌
func (bucket *TokenBucket) reFill() {
	elaps := time.Since(bucket.lastFillTime).Seconds()

	fillTokens := elaps * bucket.fillRate
	bucket.tokens += fillTokens

	if bucket.tokens > bucket.capacity {
		bucket.tokens = bucket.capacity
	}
	bucket.lastFillTime = time.Now()
}

func (bucket *TokenBucket) Allow() bool {
	bucket.mu.Lock()
	defer bucket.mu.Unlock()

	// 1. 先加入令牌
	bucket.reFill()

	// 有令牌放行
	if bucket.tokens > 0 {
		bucket.tokens -= 1
		return true
	}

	return false
}

func main() {
	// 创建一个令牌桶：桶容量为 5，每秒放入 1 个令牌
	bucket := NewTokenBucket(5, 1)

	// 模拟高频请求
	fmt.Println("开始模拟请求：")
	for i := 0; i < 10; i++ {
		if bucket.Allow() {
			fmt.Printf("请求 %d: 允许\n", i+1)
		} else {
			fmt.Printf("请求 %d: 拒绝\n", i+1)
		}
		// 每200毫秒发起一个请求
		time.Sleep(200 * time.Millisecond)
	}

	// 等待 3 秒，让令牌桶恢复
	fmt.Println("\n等待3秒后继续测试：")
	time.Sleep(3 * time.Second)

	for i := 0; i < 10; i++ {
		if bucket.Allow() {
			fmt.Printf("请求 %d: 允许\n", i+1)
		} else {
			fmt.Printf("请求 %d: 拒绝\n", i+1)
		}
		time.Sleep(200 * time.Millisecond)
	}
}
