// +build amd64
// +build !safe,!appengine

package bzero

import (
	"unsafe"
)

// asmBZero implements bzero using 16-byte / 128-bit AVX instructions; minimum size is 64 bytes.
func asmBZero(base unsafe.Pointer, size uintptr)
