//go:build !windows

package execx

// CreationFlags is a no-op on non-Windows platforms.
// @group OS Controls
//
// Example: creation flags
//
//	cmd := execx.Command("go", "env", "GOOS").CreationFlags(0)
//	fmt.Println(cmd != nil)
//	// #bool true
func (c *Cmd) CreationFlags(_ uint32) *Cmd {
	return c
}

// HideWindow is a no-op on non-Windows platforms.
// @group OS Controls
//
// Example: hide window
//
//	cmd := execx.Command("go", "env", "GOOS").HideWindow(true)
//	fmt.Println(cmd != nil)
//	// #bool true
func (c *Cmd) HideWindow(_ bool) *Cmd {
	return c
}
