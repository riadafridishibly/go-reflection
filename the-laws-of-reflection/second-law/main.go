package main

import (
	"fmt"
	"reflect"
)

func main() {
	//  Reflection goes from interface values to reflection objects and back again.
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Printf("%e\n", v.Interface()) // back in interface{}
	fmt.Println(v.Interface())
}
