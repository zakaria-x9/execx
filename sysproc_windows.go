//go:build windows

package execx

import "syscall"

// Setpgid is a no-op on Windows.
// @group OS Controls
//
// Example: setpgid
//
//	cmd := execx.Command("go", "env", "GOOS").Setpgid(true)
//	fmt.Println(cmd != nil)
//	// #bool true
func (c *Cmd) Setpgid(_ bool) *Cmd {
	return c
}

// Setsid is a no-op on Windows.
// @group OS Controls
//
// Example: setsid
//
//	cmd := execx.Command("go", "env", "GOOS").Setsid(true)
//	fmt.Println(cmd != nil)
//	// #bool true
func (c *Cmd) Setsid(_ bool) *Cmd {
	return c
}

// Pdeathsig is a no-op on Windows.
// @group OS Controls
//
// Example: pdeathsig
//
//	cmd := execx.Command("go", "env", "GOOS").Pdeathsig(0)
//	fmt.Println(cmd != nil)
//	// #bool true
func (c *Cmd) Pdeathsig(_ syscall.Signal) *Cmd {
	return c
}

// CreationFlags sets Windows creation flags.
// @group OS Controls
//
// Example: creation flags
//
//	cmd := execx.Command("go", "env", "GOOS").CreationFlags(0)
//	fmt.Println(cmd != nil)
//	// #bool true
func (c *Cmd) CreationFlags(flags uint32) *Cmd {
	c.ensureSysProcAttr()
	c.sysProcAttr.CreationFlags = flags
	return c
}

// HideWindow controls window visibility.
// @group OS Controls
//
// Example: hide window
//
//	cmd := execx.Command("go", "env", "GOOS").HideWindow(true)
//	fmt.Println(cmd != nil)
//	// #bool true
func (c *Cmd) HideWindow(on bool) *Cmd {
	c.ensureSysProcAttr()
	c.sysProcAttr.HideWindow = on
	return c
}

func (c *Cmd) ensureSysProcAttr() {
	if c.sysProcAttr == nil {
		c.sysProcAttr = &syscall.SysProcAttr{}
	}
}
