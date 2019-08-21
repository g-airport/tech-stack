package trick

import (
	"fmt"
	"reflect"
	"unsafe"
)

// another package
func GetAddFunc() interface{} {
	return add
}

type i32 int32

func  add (a, b i32) i32{
	return a + b
}

// current package
func Call() {
	add := GetAddFunc()
	addPtr := reflect.ValueOf(add).Pointer()

	// this pointer addr
	p := &addPtr
	_add := *(*func(a,b i32) int32)(unsafe.Pointer(&p))
	fmt.Println(_add(1,2))
}

