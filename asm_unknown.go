// +build !amd64
// +build !safe,!appengine

package bzero

import (
	"unsafe"
)

func asmBZero(base unsafe.Pointer, size uintptr) {
	unsafeBZero(base, size)
}
