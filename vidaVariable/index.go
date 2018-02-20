package main

import "fmt"

func main(){
	vari := hola()
	fmt.Println(vari)
	fmt.Println(*vari)
	var2 := hola()
	fmt.Println(*var2 + *vari)
}

func hola() *int{
	uno := 100
	puno := &uno

	return puno

}
