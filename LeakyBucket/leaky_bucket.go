package LeakyBucket

import (
	"sync/atomic"
	"time"
	"unsafe"
)

/********************************************************
* @author: Ihc
* @date: 2022/4/19 0019 17:21
* @version: 1.0
* @description: 漏桶算法(流量限制)
*********************************************************/

// state 上一次调用时间
type state struct {
	last  time.Time     // last 上次调用时间
	sleep time.Duration // sleep 睡眠时间
}

// LeakyBucket 漏桶
type LeakyBucket struct {
	state      unsafe.Pointer
	padding    [56]byte
	maxSlack   time.Duration
	perRequest time.Duration
}

// Take 获取请求
func (l *LeakyBucket) Take() time.Time {
	var (
		newState state
		taken    bool
		interval time.Duration
	)
	for !taken {
		now := time.Now()
		previousStatePointer := atomic.LoadPointer(&l.state)
		oldState := (*state)(previousStatePointer)
		newState = state{
			last:  now,
			sleep: oldState.sleep,
		}
		// 第一次使用Take
		if oldState.last.IsZero() {
			taken = atomic.CompareAndSwapPointer(&l.state, previousStatePointer, unsafe.Pointer(&newState))
			continue
		}
		newState.sleep += l.perRequest - now.Sub(oldState.last)
		if newState.sleep < l.maxSlack {
			newState.sleep = l.maxSlack
		}
		if newState.sleep > 0 {
			newState.last = newState.last.Add(newState.sleep)
			interval, newState.sleep = newState.sleep, 0
		}
		taken = atomic.CompareAndSwapPointer(&l.state, previousStatePointer, unsafe.Pointer(&newState))
	}
	time.Sleep(interval)
	return newState.last
}
