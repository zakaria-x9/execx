//go:build unix

package execx

import (
	"os"
	"syscall"
)

func signalFromState(state *os.ProcessState) os.Signal {
	if state == nil {
		return nil
	}
	waitStatus := state.Sys().(syscall.WaitStatus)
	if waitStatus.Signaled() {
		return waitStatus.Signal()
	}
	return nil
}
