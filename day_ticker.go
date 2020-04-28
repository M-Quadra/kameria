package kameria

import (
	"sync"
	"time"
)

// DayTicker once every day
type DayTicker struct {
	hour   int
	minute int
	second int

	rwHour   sync.RWMutex
	rwMinute sync.RWMutex
	rwSecond sync.RWMutex
}

// Hour [0, 23]
//  Get when hr is empty
//  Set when hr has value
func (slf *DayTicker) Hour(hr ...int) int {
	if len(hr) <= 0 {
		slf.rwHour.RLock()
		val := slf.hour
		slf.rwHour.RUnlock()
		return val
	}

	newVal := Limit4Int(0, hr[0], 23)
	slf.rwHour.Lock()
	slf.hour = newVal
	slf.rwHour.Unlock()
	return newVal
}

// Minute [0, 59]
//  Get when min is empty
//  Set when min has value
func (slf *DayTicker) Minute(min ...int) int {
	if len(min) <= 0 {
		slf.rwMinute.RLock()
		val := slf.minute
		slf.rwMinute.RUnlock()
		return val
	}

	newVal := Limit4Int(0, min[0], 59)
	slf.rwMinute.Lock()
	slf.minute = newVal
	slf.rwMinute.Unlock()
	return newVal
}

// Second [0, 59]
//  Get when sec is empty
//  Set when sec has value
func (slf *DayTicker) Second(sec ...int) int {
	if len(sec) <= 0 {
		slf.rwSecond.RLock()
		val := slf.second
		slf.rwSecond.RUnlock()
		return val
	}

	newVal := Limit4Int(0, sec[0], 59)
	slf.rwSecond.Lock()
	slf.second = newVal
	slf.rwSecond.Unlock()
	return newVal
}

// NewDayTicker (hour, minute, second)
func NewDayTicker(hr, min, sec int) *DayTicker {
	opt := DayTicker{}
	opt.Hour(hr)
	opt.Minute(min)
	opt.Second(sec)
	return &opt
}

// Next to next time ±1s
func (slf *DayTicker) Next() time.Time {
	needTime := slf.needTime()
	nowTime := time.Now()
	for needTime.Unix() < nowTime.Unix() {
		needTime = needTime.Add(24 * time.Hour)
	}

	len := Max4Int64(1000000000, needTime.UnixNano()-nowTime.UnixNano())
	tm := time.NewTimer(time.Duration(len))
	<-tm.C
	return needTime
}

func (slf *DayTicker) needTime() time.Time {
	t := time.Unix(time.Now().Unix(), 0).In(TimeLocation)
	t = t.Add(-time.Duration(t.Hour()) * time.Hour)
	t = t.Add(-time.Duration(t.Minute()) * time.Minute)
	t = t.Add(-time.Duration(t.Second()) * time.Second)

	t = t.Add(time.Duration(slf.Hour()) * time.Hour)
	t = t.Add(time.Duration(slf.Minute()) * time.Minute)
	t = t.Add(time.Duration(slf.Second()) * time.Second)
	return t
}
