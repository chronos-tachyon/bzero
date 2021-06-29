package bzero

// Uint8 fills the given []uint8 slice with zeroes.
func Uint8(slice []uint8) {
	dispatchUint8(slice)
}

// Uint16 fills the given []uint16 slice with zeroes.
func Uint16(slice []uint16) {
	dispatchUint16(slice)
}

// Uint32 fills the given []uint32 slice with zeroes.
func Uint32(slice []uint32) {
	dispatchUint32(slice)
}

// Uint64 fills the given []uint64 slice with zeroes.
func Uint64(slice []uint64) {
	dispatchUint64(slice)
}

// Int8 fills the given []int8 slice with zeroes.
func Int8(slice []int8) {
	dispatchInt8(slice)
}

// Int16 fills the given []int16 slice with zeroes.
func Int16(slice []int16) {
	dispatchInt16(slice)
}

// Int32 fills the given []int32 slice with zeroes.
func Int32(slice []int32) {
	dispatchInt32(slice)
}

// Int64 fills the given []int64 slice with zeroes.
func Int64(slice []int64) {
	dispatchInt64(slice)
}

func safeUint8(slice []uint8) {
	for index := range slice {
		slice[index] = 0
	}
}

func safeUint16(slice []uint16) {
	for index := range slice {
		slice[index] = 0
	}
}

func safeUint32(slice []uint32) {
	for index := range slice {
		slice[index] = 0
	}
}

func safeUint64(slice []uint64) {
	for index := range slice {
		slice[index] = 0
	}
}

func safeInt8(slice []int8) {
	for index := range slice {
		slice[index] = 0
	}
}

func safeInt16(slice []int16) {
	for index := range slice {
		slice[index] = 0
	}
}

func safeInt32(slice []int32) {
	for index := range slice {
		slice[index] = 0
	}
}

func safeInt64(slice []int64) {
	for index := range slice {
		slice[index] = 0
	}
}
