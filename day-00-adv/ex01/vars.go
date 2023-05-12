package main

import "fmt"

var i16 int16 = 42
/*
Defined at the package level, a variable is visible across all the files that
compose the package.
*/

func returnNil() interface{} {
	return nil
}

func main() {
	var str, i = "42", 42 // int -> 64 bits on a 64bit system, 32 on a 32bit system
	var ui uint
	ui = 42
	const ui8 byte = 42 // byte == uint8 ; a "string" is a set of "byte"s
	var ui32 uint32 = 42
	var i64 int64 = 42
	b, f64 := false, 42.0
	var f32 float32 = 42
	var c64 complex64 = complex(42, 0)
	c128 := complex(42, 21)
	type FortyTwo struct {
	}
	ft := FortyTwo{}
	var arr = [...]int{42} // arrays cannot be resized
	/* arrays are "value types" which means that copying them:
		arr2 := arr
	or passing an array to a function, or returning it from a function,
	creates a copy of the original array. */
	ftMap := make(map[string]int)
	ftMap["42"] = 42
	// or ftMap := map[string]int{"42":42}
	var ptr *int
	var sli = []int{}
	var ch chan int

	fmt.Printf("%s is a %T.\n", str, str)
	fmt.Printf("%d is a %T.\n", ui, ui)
	fmt.Printf("%v is an %T.\n", i, i)
	fmt.Print(ui8, " is a ")
	fmt.Printf("%T.\n", ui8)
	fmt.Printf("%d is an %T.\n", i16, i16)
	fmt.Printf("%d is a %T.\n", ui32, ui32)
	fmt.Printf("%d is an %T.\n", i64, i64)

	fmt.Printf("%t is a %T.\n", b, b)

	fmt.Printf("%g is a %T.\n", f32, f32)
	fmt.Printf("%.f is a %T.\n", f64, f64)

	fmt.Printf("%g is a %T.\n", c64, c64)
	fmt.Printf("%.f is a %T.\n", c128, c128)

	fmt.Printf("%v is a %T.\n", ft, ft)
	fmt.Printf("%v is a %T.\n", arr, arr)
	fmt.Printf("%v is a %T.\n", ftMap, ftMap)
	fmt.Printf("%p is an %T.\n", ptr, ptr) // %v gives <nil>
	fmt.Printf("%v is a %T.\n", sli, sli)
	fmt.Printf("%v is a %T.\n", ch, ch)

	fmt.Printf("%v is a %T.\n", returnNil(), returnNil())
}
