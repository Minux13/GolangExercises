package main

import "fmt"
// pc[i] is the population count of i.
var pc [256] int

func init() {
    for i := range pc {
        pc[i] = i
    }
}

// PopCount returns the population count (number of set bits) of x.
func main () {
	for i:= range pc{
		fmt.Println(pc[i])
	}
}
