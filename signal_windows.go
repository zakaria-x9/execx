//go:build windows

package execx

import "os"

func signalFromState(_ *os.ProcessState) os.Signal {
	return nil
}
