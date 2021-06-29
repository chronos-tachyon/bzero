// +build !safe,!appengine

package bzero

import (
	"reflect"
	"unsafe"
)

// Unsafe fills the given memory region with zeroes.
func Unsafe(base unsafe.Pointer, size uintptr) {
	if size < 64 {
		unsafeImpl(base, size)
		return
	}
	asmImpl(base, size)
}

func unsafeImpl(base unsafe.Pointer, size uintptr) {
	const maxInt = int(^uint(0) >> 1)
	if size > uintptr(maxInt) {
		for index := uintptr(0); index < size; index++ {
			pointer := (*uint8)(unsafe.Pointer(uintptr(base) + index))
			*pointer = 0
		}
		return
	}

	var slice []byte
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	header.Data = uintptr(base)
	header.Len = int(size)
	header.Cap = int(size)
	for index := range slice {
		slice[index] = 0
	}

}

func dispatchUint8(slice []uint8) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	Unsafe(unsafe.Pointer(header.Data), uintptr(header.Len)<<0)
}

func dispatchUint16(slice []uint16) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	Unsafe(unsafe.Pointer(header.Data), uintptr(header.Len)<<1)
}

func dispatchUint32(slice []uint32) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	Unsafe(unsafe.Pointer(header.Data), uintptr(header.Len)<<2)
}

func dispatchUint64(slice []uint64) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	Unsafe(unsafe.Pointer(header.Data), uintptr(header.Len)<<3)
}

func dispatchInt8(slice []int8) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	Unsafe(unsafe.Pointer(header.Data), uintptr(header.Len)<<0)
}

func dispatchInt16(slice []int16) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	Unsafe(unsafe.Pointer(header.Data), uintptr(header.Len)<<1)
}

func dispatchInt32(slice []int32) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	Unsafe(unsafe.Pointer(header.Data), uintptr(header.Len)<<2)
}

func dispatchInt64(slice []int64) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	Unsafe(unsafe.Pointer(header.Data), uintptr(header.Len)<<3)
}
