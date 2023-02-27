//go:build debug

package doprofiling

import (
	"context"
	"fmt"
	"os"
	"runtime/trace"
)

func SayHello() {
	fmt.Println("welcome to debug")
}

func (dp *DoProfiling) OpenTracing(filename string, topLevel string) {
	var err error
	if len(filename) == 0 {
		filename = "trace.out"
	}
	dp.tracefile, err = os.Create(filename)
	if err != nil {
		panic("unable to open trace file")
	}
	trace.Start(dp.tracefile)
	//defer
	dp.basectx = context.TODO()
	_, dp.task = trace.NewTask(dp.basectx, topLevel)
	//s.log.Debugf("Opened task, with output to file")
}

func (dp *DoProfiling) CloseTracing() {
	dp.task.End()
	trace.Stop()
	dp.tracefile.Close()
	//s.log.Debugf("Closed task, stopped trace and closed file")
}

func NewDoProfiling(filename string /*, mLogger *LeveledLogger*/) *DoProfiling {
	mDoProfiling := DoProfiling{}
	//mDoProfiling.mLogger
	mDoProfiling.OpenTracing(filename, "Request")
	return &mDoProfiling
}
