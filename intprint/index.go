package main

import (
    //"math/rand"
    "fmt"
)

func main() {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
	b:=3.134416832398
	fmt.Printf("%f %e  %g", b, b, b)
	// Output:
	// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF
}
