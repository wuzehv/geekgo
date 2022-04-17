package main

import (
	"fmt"
	"sync"
	"time"
)

// rollingNumber 数字载体，用简单类型
type rollingNumber int64

// rollingNumberContainer 整个滑动窗口的计数容器，采用map结构
type rollingNumberContainer map[int64]*rollingNumber

// container 滑动窗口，均使用非导出字段
type container struct {
	lock sync.RWMutex
	// size 滑动窗口中统计的数量；每秒一个容器
	size int64
	rn   rollingNumberContainer
}

// getNumber 获取底层计数
func (h *container) getNumber() *rollingNumber {
	key := time.Now().Unix()
	if _, ok := h.rn[key]; !ok {
		var tmp rollingNumber
		h.rn[key] = &tmp
	}

	return h.rn[key]
}

// setNumber 设置数字
func (n *rollingNumber) setNumber(num int64) {
	*n += rollingNumber(num)
}

// removeItem 移除过期的容器
func (h *container) removeItem() {
	diff := time.Now().Unix() - h.size
	for k := range h.rn {
		if k <= diff {
			delete(h.rn, k)
		}
	}
}

// Incr 增加
func (h *container) Incr(t int64) {
	h.lock.Lock()
	defer h.lock.Unlock()
	x := h.getNumber()
	x.setNumber(t)
	h.removeItem()
}

func NewContainer(size uint) *container {
	return &container{size: int64(size), rn: make(rollingNumberContainer, 0)}
}

func main() {
	x := NewContainer(10)

	go func() {
		for {
			x.Incr(1)
			time.Sleep(1 * time.Millisecond)
		}
	}()

	go func() {
		for {
			x.Incr(1)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	time.Sleep(3 * time.Second)
	for k, v := range x.rn {
		fmt.Printf("key: %v, number: %d\n", k, *v)
	}
}
