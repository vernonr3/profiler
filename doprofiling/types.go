package doprofiling

import (
	"context"
	"os"
	"runtime/trace"
)

type DoProfiling struct {
	tracefile *os.File
	task      *trace.Task
	basectx   context.Context
	// mLogger   *LeveledLogger
}
