package crypt

import (
	"reflect"
	"unsafe"
)

// string could really be whatever
func MemClear(s *string) {
	if s == nil {
		return
	}

	// deprecated, and incredibly unsafe, but that is the cost of memory clearing
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(s))
	data := (*[1 << 30]byte)(unsafe.Pointer(stringHeader.Data))[:stringHeader.Len:stringHeader.Len]

	for i := range data {
		data[i] = 1
	}

}
