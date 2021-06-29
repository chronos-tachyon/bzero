// +build amd64
// +build !safe,!appengine

package bzero

import (
	"unsafe"
)

// asmImpl implements bzero using 16-byte / 128-bit AVX instructions; minimum size is 64 bytes.
func asmImpl(base unsafe.Pointer, size uintptr)
