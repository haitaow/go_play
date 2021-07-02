package main

import (
	"fmt"
	"strconv"
)

type I1 interface {
	Foo() int
}

type I2 interface {
	Foo() string
}

type X struct {
	value int
}

func (x X) Foo() int {
	return x.value
}

// Foo2 implements I2， but cannot use name Foo again
func (x X) Foo2() string {
	return strconv.Itoa(x.value)
}

func printI1(x I1) {
	v := x.Foo()
	fmt.Printf("%#v(%T)\n", v, v)
}

func printI2(x I2) {
	v := x.Foo()
	fmt.Printf("%#v(%T)\n", v, v)
}

func main() {
	x := &X{101}

	printI1(x)

	// compile error; how to make this work?
	// printI2(x)

	workaround(x) // scroll down to see
}

// XAsI2 is a workaround for interface method name clash and essentially implements I2 with X.Foo2.
// This can be a generic pattern.
type XAsI2 struct{ X *X }

func (x XAsI2) Foo() string {
	return x.X.Foo2()
}

func workaround(x *X) {
	printI2(XAsI2{x})
}
