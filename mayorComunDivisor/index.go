package main

import ("fmt"
		"bufio"
		"os"
		"strings"
		"strconv"
		)

func main(){
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		texto := input.Text()
		te := strings.Split(texto, " ")
		if texto != ""{
		a, err1 := strconv.ParseInt(te[0],10,0)
		b, err2 := strconv.ParseInt(te[1],10,0)
		if err1 != nil && err2 != nil{
			fmt.Println("Error")
		}else{
			a, b = b, a%b
			fmt.Println(a, b)
			}
		}
	}

}
