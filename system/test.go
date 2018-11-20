package system

import (
	"fmt"
	"time"
)

//TestTime time test func
func TestTime() {
	t := NewSysTime()
	c := make(chan int)
	go (func() {
		for i := 0; i < 100; i++ {
			fmt.Printf("TestTime system time:%v\n", t.GetGlobalTimeMS())
			time.Sleep(time.Duration(10*i) * time.Millisecond)
		}
		c <- 1
	})()

	<-c
}
