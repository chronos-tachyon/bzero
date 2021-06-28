// +build !safe,!appengine

package bzero

import (
	"reflect"
	"unsafe"
)

// BZeroUnsafe fills the given memory region with zeroes.
func BZeroUnsafe(base unsafe.Pointer, size uintptr) {
	if size < 64 {
		unsafeBZero(base, size)
		return
	}
	asmBZero(base, size)
}

func unsafeBZero(base unsafe.Pointer, size uintptr) {
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

func dispatchBZeroUint8(slice []uint8) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	BZeroUnsafe(unsafe.Pointer(header.Data), uintptr(header.Len)<<0)
}

func dispatchBZeroUint16(slice []uint16) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	BZeroUnsafe(unsafe.Pointer(header.Data), uintptr(header.Len)<<1)
}

func dispatchBZeroUint32(slice []uint32) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	BZeroUnsafe(unsafe.Pointer(header.Data), uintptr(header.Len)<<2)
}

func dispatchBZeroUint64(slice []uint64) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	BZeroUnsafe(unsafe.Pointer(header.Data), uintptr(header.Len)<<3)
}

func dispatchBZeroInt8(slice []int8) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	BZeroUnsafe(unsafe.Pointer(header.Data), uintptr(header.Len)<<0)
}

func dispatchBZeroInt16(slice []int16) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	BZeroUnsafe(unsafe.Pointer(header.Data), uintptr(header.Len)<<1)
}

func dispatchBZeroInt32(slice []int32) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	BZeroUnsafe(unsafe.Pointer(header.Data), uintptr(header.Len)<<2)
}

func dispatchBZeroInt64(slice []int64) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	BZeroUnsafe(unsafe.Pointer(header.Data), uintptr(header.Len)<<3)
}
