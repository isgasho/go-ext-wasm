package wasmer

import (
	"reflect"
	"unsafe"
)

type Memory struct {
	data   *uint8
	length uint32
}

func NewMemory(data *uint8, length uint32) Memory {
	return Memory{data, length}
}

func (self *Memory) GetLength() uint32 {
	return self.length
}

func (self *Memory) GetData() []uint8 {
	var raw []uint8
	var header reflect.SliceHeader = *(*reflect.SliceHeader)(unsafe.Pointer(&raw))

	header.Data = uintptr(unsafe.Pointer(self.data))
	header.Len = int(self.length)
	header.Cap = int(self.length)

	return *(*[]uint8)(unsafe.Pointer(&header))
}
