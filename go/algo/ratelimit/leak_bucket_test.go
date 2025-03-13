package main

import (
	"fmt"
	"testing"
	"time"
)

func TestLeakBucekt(t *testing.T) {

	// 创建一个容量为5，每200毫秒处理一个请求的漏桶
	bucket := NewLeakBucket(5, 200*time.Millisecond)

	// 模拟请求
	for i := 1; i <= 10; i++ {
		// 尝试加入请求
		if bucket.Allow() {
			fmt.Printf("请求 %d 成功加入漏桶，等待处理\n", i)
		} else {
			fmt.Printf("请求 %d 被拒绝，漏桶已满\n", i)
		}

		// 模拟请求间隔
		if i == 5 {
			fmt.Println("等待1秒...")
			time.Sleep(1 * time.Second)
		} else {
			time.Sleep(100 * time.Millisecond)
		}
	}

	// 等待所有请求处理完成
	time.Sleep(2 * time.Second)
	bucket.Leak()
	fmt.Printf("处理完成后，桶中剩余水量: %d\n", bucket.current)

}
