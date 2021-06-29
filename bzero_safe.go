// +build safe appengine

package bzero

func dispatchUint8(slice []uint8) {
	safeUint8(slice)
}

func dispatchUint16(slice []uint16) {
	safeUint16(slice)
}

func dispatchUint32(slice []uint32) {
	safeUint32(slice)
}

func dispatchUint64(slice []uint64) {
	safeUint64(slice)
}

func dispatchInt8(slice []int8) {
	safeInt8(slice)
}

func dispatchInt16(slice []int16) {
	safeInt16(slice)
}

func dispatchInt32(slice []int32) {
	safeInt32(slice)
}

func dispatchInt64(slice []int64) {
	safeInt64(slice)
}
