package main

import (
	"sync"
	"time"
)

type LeakBucket struct {
	capacity int           // 容量
	current  int           // 当前水量
	leakRate time.Duration // 楼桶速率
	lastLeak time.Time     // 上次楼桶时间
	mu       sync.Mutex
}

// 楼桶的原理是：
//  1. 固定数量的桶，以恒定的速率漏水，请求时进水
//  2. 请求时判读：
//     2.1 先漏水
//     2.2 再判断，当前水量是否小于固定容量，是，则放行；否则，限流
//     2.3 进水（当前水量++）
func NewLeakBucket(cap int, leakRate time.Duration) *LeakBucket {
	return &LeakBucket{
		capacity: cap,
		current:  0,
		leakRate: leakRate,
		lastLeak: time.Now(),
		mu:       sync.Mutex{},
	}
}

// 漏桶：
// 1. 更新当前水量
// 2. 更新上次楼桶时间
func (lb *LeakBucket) Leak() {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	elaps := time.Since(lb.lastLeak).Seconds()
	leakAmt := int(elaps / float64(lb.leakRate))
	lb.lastLeak = time.Now()

	lb.current -= leakAmt
	if lb.current < 0 {
		lb.current = 0
	}
}

func (lb *LeakBucket) Allow() bool {
	// 1. 先漏水
	lb.Leak()

	lb.mu.Lock()
	defer lb.mu.Unlock()

	if lb.current > lb.capacity {
		return false
	}

	lb.current++
	return true
}
