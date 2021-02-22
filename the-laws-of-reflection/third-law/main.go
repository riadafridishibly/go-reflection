package main

import (
	"fmt"
	"reflect"
)

func main() {
	// To modify a reflection object, the value must be settable.
	var x float64 = 3.4
	v := reflect.ValueOf(x) // passing by value here
	// v.SetFloat(7.1) // panic:

	// Settability is a property of a reflection `Value`

	fmt.Println("settability of v:", v.CanSet()) // settability of v: false

	// let's try setting it
	p := reflect.ValueOf(&x)                     // taking the address
	fmt.Println("type of p:", p.Type())          // type of p: *float64
	fmt.Println("settability of p:", p.CanSet()) // settability of p: false

	vp := p.Elem()
	vi := reflect.Indirect(p)
	fmt.Println("p.Elem() == reflect.Indirect(p):", vp == vi)
	fmt.Println("settability of vp:", vp.CanSet()) // settability of vp: true
	fmt.Println("settability of vi:", vi.CanSet()) // settability of vi: true

	// structs
	type T struct {
		A int
		B string
		// p rune
	}

	t := T{23, "skidoo"}
	s := reflect.Indirect(reflect.ValueOf(&t))
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v ", i, typeOfT.Field(i).Name, f.Type(), f.Interface()) // f.Interface() call will fail for `p rune` case
		fmt.Printf("is settable: %v, is addressable: %v\n", f.CanSet(), f.CanAddr())
	}

	fmt.Printf("before: %+v\n", t) // before: {A:23 B:skidoo}

	s.Field(0).SetInt(77)
	s.Field(1).SetString("Euphoria")

	fmt.Printf("after: %+v\n", t) // after: {A:77 B:Euphoria}
}
