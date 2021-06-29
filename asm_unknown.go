// +build !amd64
// +build !safe,!appengine

package bzero

import (
	"unsafe"
)

func asmImpl(base unsafe.Pointer, size uintptr) {
	unsafeImpl(base, size)
}
