package ghostferry

import (
	"sync"
	"sync/atomic"
	"time"
)

type Throttler interface {
	Throttled() bool
	SetPaused(bool)
	Run(*sync.WaitGroup)
	Stop()
}

func WaitForThrottle(t Throttler) {
	for {
		if !t.Throttled() {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
}

type PauserThrottler struct {
	paused int32
}

func (t *PauserThrottler) Throttled() bool {
	return atomic.LoadInt32(&t.paused) != 0
}

func (t *PauserThrottler) SetPaused(paused bool) {
	var val int32
	if paused {
		val = 1
	}
	atomic.StoreInt32(&t.paused, val)
}

func (t *PauserThrottler) Run(wg *sync.WaitGroup) {
	wg.Done()
}

func (t *PauserThrottler) Stop() {}
