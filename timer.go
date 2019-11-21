package kameria

import (
	"sync"
	"time"
)

//DayTimer default no repeat
type DayTimer struct {
	Hour     int  //[0, 23]
	Minute   int  //[0, 59]
	Second   int  //[0, 59]
	IsRepeat bool //default false
	Handler  func()

	wannaCancel bool
	ticker      *time.Ticker
	fstTimer    *time.Timer
	mutex       sync.RWMutex
}

func (slf *DayTimer) getWannaCancel() bool {
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()

	return slf.wannaCancel
}

func (slf *DayTimer) setWannaCancel(v bool) {
	slf.mutex.Lock()
	defer slf.mutex.Unlock()

	slf.wannaCancel = v
}

//Start once
func (slf *DayTimer) Start() {
	slf.mutex.Lock()
	defer slf.mutex.Unlock()

	slf.wannaCancel = false
	if slf.ticker == nil {
		// slf.ticker = time.NewTicker(24 * time.Hour)
		slf.ticker = time.NewTicker(1 * time.Second)
	}
	if slf.fstTimer != nil {
		return
	}

	if slf.Handler == nil {
		return
	}

	needTime := time.Now()
	needTime = needTime.Add(-time.Duration(needTime.Hour()) * time.Hour)
	needTime = needTime.Add(-time.Duration(needTime.Minute()) * time.Minute)
	needTime = needTime.Add(-time.Duration(needTime.Second()) * time.Second)
	needTime = needTime.Add(-time.Duration(needTime.Nanosecond()) * time.Nanosecond)

	slf.Hour = Limit4Int(0, slf.Hour, 23)
	slf.Minute = Limit4Int(0, slf.Minute, 59)
	slf.Second = Limit4Int(0, slf.Second, 59)

	needTime = needTime.Add(time.Duration(slf.Hour) * time.Hour)
	needTime = needTime.Add(time.Duration(slf.Minute) * time.Minute)
	needTime = needTime.Add(time.Duration(slf.Second) * time.Second)

	nowTime := time.Now()
	if nowTime.UnixNano() > needTime.UnixNano() {
		if !slf.IsRepeat { //out of time
			return
		}

		needTime = needTime.Add(24 * time.Hour)
	}

	len := Max4Int64(0, needTime.UnixNano()-nowTime.UnixNano())
	slf.fstTimer = time.NewTimer(time.Duration(len))

	go func() {
		<-slf.fstTimer.C
		if slf.getWannaCancel() {
			slf.setWannaCancel(false)
			slf.fstTimer = nil
			return
		}

		go slf.Handler()
		if !slf.IsRepeat {
			slf.fstTimer = nil
			return
		}

		go slf.again()
	}()
}

func (slf *DayTimer) again() {
	for range slf.ticker.C {
		if slf.getWannaCancel() {
			slf.setWannaCancel(false)
			slf.fstTimer = nil
			return
		}
		if slf.Handler == nil || !slf.IsRepeat {
			slf.fstTimer = nil
			return
		}

		go slf.Handler()
	}

}

//Cancel timer
func (slf *DayTimer) Cancel() {
	slf.setWannaCancel(true)
}
