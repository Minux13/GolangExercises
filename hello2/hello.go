package main
import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i:=1; i<len(os.Args); i++{
		s += sep + os.Args[i]
		sep = " "
		fmt.Println(len(os.Args[i]))
	}

	fmt.Println(len(os.Args))
	fmt.Println(os.Args[0])
}
