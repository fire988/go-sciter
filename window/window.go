package window

import (
	"runtime"

	"github.com/fire988/go-sciter"
)

type Window struct {
	*sciter.Sciter
	creationFlags sciter.WindowCreationFlag
}

func (w *Window) run() {
	// runtime.LockOSThread()
}

func init() {
	runtime.LockOSThread()
}
