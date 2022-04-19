package TokenBucket

import (
	"math"
	"strconv"
	"sync"
	"time"
)

/****************************************************************
 * @author: Ihc
 * @date: 2022/4/19 19:42
 * @description: 令牌桶
 ***************************************************************/

const infinityDuration time.Duration = 0x7fffffffffffffff

// Clock 时钟接口
type Clock interface {
	// Now 获取时钟的当前时间
	Now() time.Time
	// Sleep 睡眠规定的时间
	Sleep(duration time.Duration)
}

// realClock 用标准库时间模块实现Clock接口
type realClock struct{}

func (r realClock) Now() time.Time {
	return time.Now()
}

func (r realClock) Sleep(duration time.Duration) {
	time.Sleep(duration)
}

// TokenBucket 令牌桶
type TokenBucket struct {
	clock           Clock         // clock
	capacity        int64         // capacity 桶的容量
	startTime       time.Time     // startTime 令牌桶被创建时的时间
	quantum         int64         // quantum 每个滴答加入多少token
	interval        time.Duration // interval 滴答的间隔时间
	mu              sync.Mutex    // mu 用于线程同步
	availableTokens int64         // availableTokens 可用的令牌数量
	latestTick      int64         // latestTick 最近一次的滴答时间点
}

// nextQuantum
func nextQuantum(n int64) int64 {
	m := n * 11 / 10
	if m == n {
		m++
	}
	return m
}

// currentTick 当前的滴答数量
// 当前滴答数量 = 令牌桶创建时间到指定时钟的当前时间的时间间隔 / 每个滴答的时间间隔
func (t *TokenBucket) currentTick(now time.Time) int64 {
	return int64(now.Sub(t.startTime) / t.interval)
}

// available 刷新桶内令牌数量
func (t *TokenBucket) available(now time.Time) int64 {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.adjustAvailableTokens(t.currentTick(now))
	return t.availableTokens
}

// adjustAvailableTokens 增加桶内的令牌数量
// tick表示当前滴答的时间点
func (t *TokenBucket) adjustAvailableTokens(tick int64) {
	lastTick := t.latestTick
	t.latestTick = tick
	if t.availableTokens >= t.capacity {
		return
	}
	t.availableTokens += (tick - lastTick) * t.quantum
	if t.availableTokens > t.capacity {
		t.availableTokens = t.capacity
	}
	return
}

// take 取出令牌
// now 表示当前时间; count 表示要取出令牌的数量; maxWait 最长等待时间
// 若超过maxWait未取得规定数量的令牌,则返回0
// 否则返回成功取得令牌数量所消耗的时间
// 返回的布尔值表示是否取得成功
func (t *TokenBucket) take(now time.Time, count int64, maxWait time.Duration) (time.Duration, bool) {
	if count <= 0 {
		return 0, true
	}
	// 根据滴答数更新桶内令牌的数量
	tick := t.currentTick(now)
	t.adjustAvailableTokens(tick)
	// 当前桶内令牌的数量足够
	available := t.availableTokens - count
	if available >= 0 {
		t.availableTokens = available
		return 0, true
	}
	// 当前桶内令牌的数量不足够
	// 计算需要的总tick,计算从now时间起到获取count数量的token需要等待的时间
	endTick := tick + (-available+t.quantum-1)/t.quantum
	endTime := t.startTime.Add(time.Duration(endTick) * t.interval)
	waitTime := endTime.Sub(now)
	// 指定的等待时间内完成规定数量的取出
	if waitTime > maxWait {
		return 0, false
	}
	// 指定的等待时间内未完成规定数量的取出,需要等待
	t.availableTokens = available
	return waitTime, true
}

// takeAvailable
func (t *TokenBucket) takeAvailable(now time.Time, count int64) int64 {
	if count <= 0 {
		return 0
	}
	t.adjustAvailableTokens(t.currentTick(now))
	if t.availableTokens <= 0 {
		return 0
	}
	if count > t.availableTokens {
		count = t.availableTokens
	}
	t.availableTokens -= count
	return count
}

// Take 取出指定数量的令牌
// 返回需要等待的时长
func (t *TokenBucket) Take(count int64) time.Duration {
	duration, _ := t.take(t.clock.Now(), count, infinityDuration)
	return duration
}

// TakeMaxDuration 取出指定数量的令牌
// 若超过maxWait时间,依然未满足要求,则返回要等待的时间和未完成标志false
// 若完成要求,则返回0和true
func (t *TokenBucket) TakeMaxDuration(count int64, maxWait time.Duration) (time.Duration, bool) {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.take(t.clock.Now(), count, maxWait)
}

// TakeAvailable 指定数量的令牌能否可用
// 返回可用的令牌数量,返回的数量最大为count
// 若无令牌可以用,则返回0
func (t *TokenBucket) TakeAvailable(count int64) int64 {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.takeAvailable(t.clock.Now(), count)
}

// Available 刷新桶内令牌数量,时间从给定时钟的当前时刻起
func (t *TokenBucket) Available() int64 {
	return t.available(t.clock.Now())
}

// Wait 取出指定数量的token,若token数量不足,则等待至token数量满足要求
func (t *TokenBucket) Wait(count int64) {
	if d := t.Take(count); d > 0 {
		t.clock.Sleep(d)
	}
}

// WaitMaxDuration 取出指定数量的令牌
// 若在指定的时间内能取出,则进行阻塞等待
// 否则直接返回
func (t *TokenBucket) WaitMaxDuration(count int64, maxWait time.Duration) bool {
	duration, ok := t.TakeMaxDuration(count, maxWait)
	if duration > 0 {
		t.clock.Sleep(duration)
	}
	return ok
}

// Rate 返回令牌桶的速率
// 速率(单位秒)
func (t *TokenBucket) Rate() float64 {
	return 1e9 * float64(t.quantum) / float64(t.interval)
}

// Capacity 返回令牌桶的容量
func (t *TokenBucket) Capacity() int64 {
	return t.capacity
}

// NewTokenBucket 创建令牌桶
// interval 表示每个滴答的时间间隔
// capacity 表示令牌桶的容量
func NewTokenBucket(interval time.Duration, capacity int64) *TokenBucket {
	return NewTokenBucketWithClock(interval, capacity, nil)
}

// NewTokenBucketWithClock 指定滴答的时间间隔,容量和时钟创建令牌桶
func NewTokenBucketWithClock(interval time.Duration, capacity int64, clock Clock) *TokenBucket {
	return NewTokenBucketWithQuantumAndClock(interval, capacity, 1, clock)
}

// NewTokenBucketWithRate 指定速率和容量创建令牌桶
func NewTokenBucketWithRate(rate float64, capacity int64) *TokenBucket {
	return NewTokenBucketWithRateAndClock(rate, capacity, nil)
}

// NewTokenBucketWithRateAndClock 指定速率,容量和时钟创建令牌桶
func NewTokenBucketWithRateAndClock(rate float64, capacity int64, clock Clock) *TokenBucket {
	t := NewTokenBucketWithQuantumAndClock(1, capacity, 1, clock)
	for quantum := int64(1); quantum < 1<<50; quantum = nextQuantum(quantum) {
		interval := time.Duration(1e9 * float64(quantum) / rate)
		if interval <= 0 {
			continue
		}
		t.interval = interval
		t.quantum = quantum
		if abs := math.Abs(t.Rate() - rate); abs/rate <= 0.01 {
			return t
		}
	}
	panic("cannot find suitable quantum for " + strconv.FormatFloat(rate, 'g', -1, 64))
}

func NewTokenBucketWithQuantumAndClock(interval time.Duration, capacity int64, quantum int64, clock Clock) *TokenBucket {
	if clock == nil {
		clock = realClock{}
	}
	if interval <= 0 {
		panic("token bucket interval is not > 0")
	}
	if quantum <= 0 {
		panic("token bucket quantum is not > 0")
	}
	return &TokenBucket{
		clock:           clock,
		startTime:       clock.Now(),
		latestTick:      0,
		interval:        interval,
		capacity:        capacity,
		quantum:         quantum,
		availableTokens: capacity,
	}
}
