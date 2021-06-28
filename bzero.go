package bzero

// BZeroUint8 fills the given []uint8 slice with zeroes.
func BZeroUint8(slice []uint8) {
	dispatchBZeroUint8(slice)
}

// BZeroUint16 fills the given []uint16 slice with zeroes.
func BZeroUint16(slice []uint16) {
	dispatchBZeroUint16(slice)
}

// BZeroUint32 fills the given []uint32 slice with zeroes.
func BZeroUint32(slice []uint32) {
	dispatchBZeroUint32(slice)
}

// BZeroUint64 fills the given []uint64 slice with zeroes.
func BZeroUint64(slice []uint64) {
	dispatchBZeroUint64(slice)
}

// BZeroInt8 fills the given []int8 slice with zeroes.
func BZeroInt8(slice []int8) {
	dispatchBZeroInt8(slice)
}

// BZeroInt16 fills the given []int16 slice with zeroes.
func BZeroInt16(slice []int16) {
	dispatchBZeroInt16(slice)
}

// BZeroInt32 fills the given []int32 slice with zeroes.
func BZeroInt32(slice []int32) {
	dispatchBZeroInt32(slice)
}

// BZeroInt64 fills the given []int64 slice with zeroes.
func BZeroInt64(slice []int64) {
	dispatchBZeroInt64(slice)
}

func safeBZeroUint8(slice []uint8) {
	for index := range slice {
		slice[index] = 0
	}
}

func safeBZeroUint16(slice []uint16) {
	for index := range slice {
		slice[index] = 0
	}
}

func safeBZeroUint32(slice []uint32) {
	for index := range slice {
		slice[index] = 0
	}
}

func safeBZeroUint64(slice []uint64) {
	for index := range slice {
		slice[index] = 0
	}
}

func safeBZeroInt8(slice []int8) {
	for index := range slice {
		slice[index] = 0
	}
}

func safeBZeroInt16(slice []int16) {
	for index := range slice {
		slice[index] = 0
	}
}

func safeBZeroInt32(slice []int32) {
	for index := range slice {
		slice[index] = 0
	}
}

func safeBZeroInt64(slice []int64) {
	for index := range slice {
		slice[index] = 0
	}
}
