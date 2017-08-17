package demo

import "testing"

func TestFoo(t *testing.T) {
	testFoo(&Foo{
		Slice: []*Bar{
			&Bar{IntVal: 1, FloatVal: 2, StringVal: "hello"},
			&Bar{IntVal: 3, FloatVal: 4, StringVal: "world"},
			nil,
		},
	})
}
