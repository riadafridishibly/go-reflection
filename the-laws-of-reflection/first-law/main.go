package main

import (
	"fmt"
	"reflect"
	"strings"
)

func horizontalLine(n int) {
	fmt.Println(strings.Repeat("-+-", n))
}

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", reflect.TypeOf(x)) // type: float64
	fmt.Println("value:", v.String())       // value: <float64 Value>

	fmt.Println("kind is float64:",
		v.Kind() == reflect.Float64) // kind is float64: true
	fmt.Println("value:", v.Float()) // value: 3.4

	horizontalLine(10)
	// the "getter" and "setter" methods of Value operate on
	// the largest type that can hold the value

	var y uint8 = 'x'
	u := reflect.ValueOf(y)
	fmt.Println("type:", u.Type())                           // type: uint8
	fmt.Println("kind is uint8:", u.Kind() == reflect.Uint8) // kind is uint8: true
	y = uint8(u.Uint())

	horizontalLine(10)
	// Kind of a reflection object describes the underlying type, not the static type.
	type MyInt int
	var z MyInt = 7
	zr := reflect.ValueOf(z)

	fmt.Println("type:", zr.Type()) // type: main.MyInt
	fmt.Println("kind:", zr.Kind()) // kind: int
}
