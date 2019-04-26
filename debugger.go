// debugger is a single step debugger
package debugger

import (
	"sync"
)

var breakpoints map[string]*BreakPoint
var wg sync.Mutex

func init() {
	breakpoints = make(map[string]*BreakPoint)
}

type BreakPoint struct {
	c chan struct{}
}

func Bind(name string) *BreakPoint {
	return getBreakPoint(name)
}

func Break(b *BreakPoint) {
	<-b.c
}

func Continue(name string) {
	b := getBreakPoint(name)
	b.c <- struct{}{}
}

func getBreakPoint(name string) *BreakPoint {
	var ok bool
	var b *BreakPoint
	wg.Lock()
	b, ok = breakpoints[name]
	if !ok {
		b = &BreakPoint{
			c: make(chan struct{}),
		}
		breakpoints[name] = b
	}
	wg.Unlock()
	return b
}
