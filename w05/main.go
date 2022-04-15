package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

const (
	// vLen 滑动窗口统计保留数量
	vLen = 10
)

// diff 为当前时间戳与长度的取模，用来保证数组下标从0开始
var diff = time.Now().Unix() % vLen

type rollingNumber [vLen]int64

func (h *rollingNumber) getKey(t int64) int64 {
	return (t - diff) % vLen
}

func (h *rollingNumber) Incr(t int64) {
	atomic.AddInt64(&h[h.getKey(time.Now().Unix())], t)
}

type HystrixContainer struct {
	all       *rollingNumber
	successes *rollingNumber
	failures  *rollingNumber
	timeouts  *rollingNumber
}


func main() {
	x := &rollingNumber{}

	go func() {
		for {
			x.Incr(1)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	time.Sleep(25 * time.Second)
	fmt.Println(x)
}
