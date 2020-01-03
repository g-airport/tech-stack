package trick

import (
	"reflect"
	"unsafe"
)

// By unsafe pointer

type Data struct {
	Int32  int32
	String string
}

var sizeofUnsafeData = int(unsafe.Sizeof(Data{}))

func Data2Bytes(s *Data) []byte {
	var x reflect.SliceHeader
	x.Len = sizeofUnsafeData
	x.Cap = sizeofUnsafeData
	x.Data = uintptr(unsafe.Pointer(s))
	return *(*[]byte)(unsafe.Pointer(&x))
}

func Bytes2Data(b []byte) *Data {
	return (*Data)(unsafe.Pointer(
		(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
	))
}



// convert b to string without copy
func Bytes2StringFast(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func String2ByteFast(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{Data: sh.Data, Len: sh.Len, Cap: 0}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func String2ByteFastV2(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{Data: sh.Data, Len: sh.Len, Cap: sh.Len}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func String2Byte(s string) []byte {
	return []byte(s)
}

func Byte2String(b []byte) string {
	return string(b)
}



// returns &s[0], which is not allowed in go
func stringPointer(s string) unsafe.Pointer {
	p := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return unsafe.Pointer(p.Data)
}

// returns &b[0], which is not allowed in go
func bytePointer(b []byte) unsafe.Pointer {
	p := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	return unsafe.Pointer(p.Data)
}


