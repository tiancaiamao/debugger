// debugger is a single step debugger
package debugger

import (
	"sync"
)

var breakpoints map[string]*Label
var wg sync.Mutex

func init() {
	breakpoints = make(map[string]*Label)
}

type Label struct {
	c chan struct{}
}

func Bind(name string) *Label {
	return getLabel(name)
}

func Breakpoint(b *Label) {
	<-b.c
}

func Continue(name string) {
	b := getLabel(name)
	b.c <- struct{}{}
}

func getLabel(name string) *Label {
	var ok bool
	var b *Label
	wg.Lock()
	b, ok = breakpoints[name]
	if !ok {
		b = &Label{
			c: make(chan struct{}),
		}
		breakpoints[name] = b
	}
	wg.Unlock()
	return b
}
