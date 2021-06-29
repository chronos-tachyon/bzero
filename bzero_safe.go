// +build safe appengine

package bzero

func dispatchUint8(slice []uint8) {
	for index := range slice {
		slice[index] = 0
	}
}

func dispatchUint16(slice []uint16) {
	for index := range slice {
		slice[index] = 0
	}
}

func dispatchUint32(slice []uint32) {
	for index := range slice {
		slice[index] = 0
	}
}

func dispatchUint64(slice []uint64) {
	for index := range slice {
		slice[index] = 0
	}
}

func dispatchInt8(slice []int8) {
	for index := range slice {
		slice[index] = 0
	}
}

func dispatchInt16(slice []int16) {
	for index := range slice {
		slice[index] = 0
	}
}

func dispatchInt32(slice []int32) {
	for index := range slice {
		slice[index] = 0
	}
}

func dispatchInt64(slice []int64) {
	for index := range slice {
		slice[index] = 0
	}
}
