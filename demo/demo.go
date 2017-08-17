package demo

// #cgo CXXFLAGS: -std=c++11
//
// #include <stdint.h>
// void TestFoo(uintptr_t foo);
import "C"
import "unsafe"

type Foo struct {
	Slice []*Bar
	Array [4]Bar
}

type Bar struct {
	IntVal    int
	FloatVal  float64
	StringVal string
}

func testFoo(foo *Foo) {
	C.TestFoo(C.uintptr_t(uintptr(unsafe.Pointer(foo))))
}
