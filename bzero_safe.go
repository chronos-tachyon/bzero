// +build safe appengine

package bzero

func dispatchBZeroUint8(slice []uint8) {
	safeBZeroUint8(slice)
}

func dispatchBZeroUint16(slice []uint16) {
	safeBZeroUint16(slice)
}

func dispatchBZeroUint32(slice []uint32) {
	safeBZeroUint32(slice)
}

func dispatchBZeroUint64(slice []uint64) {
	safeBZeroUint64(slice)
}

func dispatchBZeroInt8(slice []int8) {
	safeBZeroInt8(slice)
}

func dispatchBZeroInt16(slice []int16) {
	safeBZeroInt16(slice)
}

func dispatchBZeroInt32(slice []int32) {
	safeBZeroInt32(slice)
}

func dispatchBZeroInt64(slice []int64) {
	safeBZeroInt64(slice)
}
