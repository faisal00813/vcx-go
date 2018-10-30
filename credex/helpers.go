package credex

import (
	"sync/atomic"

	vcx "github.com/faisal00813/vcx-go/vcx"
)

type count32 int32

var counter count32

func (c *count32) increment() int32 {
	return atomic.AddInt32((*int32)(c), 1)
}
func (c *count32) get() int32 {
	return atomic.LoadInt32((*int32)(c))
}

func createCommandHandle() vcx.Command_handle_t {
	_counter := counter.increment()
	return vcx.Command_handle_t(_counter)
}
