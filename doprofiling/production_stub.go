//go:build !debug

package doprofiling

import "fmt"

func SayHello() {
	fmt.Println("welcome to production")
}

func NewDoProfiling(filename string) *DoProfiling {
	return &DoProfiling{}
}

func (dp *DoProfiling) OpenTracing(filename string, topLevel string) {}

func (dp *DoProfiling) CloseTracing() {}
