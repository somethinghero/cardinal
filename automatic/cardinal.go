package automatic

import (
	"github.com/somethinghero/cardinal/system"
)

type Cardinal struct {
	sysTime *system.SysTime
}

var (
	cardinal *Cardinal
)

func init() {
	cardinal := &Cardinal{}
	cardinal.sysTime = system.NewSysTime()
}
