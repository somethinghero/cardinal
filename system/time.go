package system

import (
	"time"
)

//SysTime system time struct
type SysTime struct {
	timeMS           int64
	lastUpdateTimeMS int64
}

//GlobalTime global system time
var (
	GlobalTime *SysTime
)

//NewSysTime new a SystemTime
func NewSysTime() *SysTime {
	ret := &SysTime{}
	ret.initTime()
	return ret
}

func init() {
	GlobalTime = NewSysTime()
}

func (t *SysTime) initTime() {
	t.lastUpdateTimeMS = time.Now().UnixNano() / 1000000
	t.timeMS = 0
}

func (t *SysTime) updateGlobalTime() {
	if t.lastUpdateTimeMS == 0 {
		t.lastUpdateTimeMS = time.Now().UnixNano() / 1000000
		t.timeMS = 0
		return
	}
	nowMS := time.Now().UnixNano() / 1000000
	t.timeMS += nowMS - t.lastUpdateTimeMS
	if t.timeMS < 0 {
		t.timeMS = 0
	}
	t.lastUpdateTimeMS = nowMS
}

//GetGlobalTimeMS get time from system start running, in ms
func (t *SysTime) GetGlobalTimeMS() int64 {
	t.updateGlobalTime()
	return t.timeMS
}
