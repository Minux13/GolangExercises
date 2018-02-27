
package main

import ("fmt"
       	"time"
      	)
var texto string = "Hola, 世界"

type Employee struct {
    ID        int
    Name      string
    Address   string
    DoB       time.Time
    Position  string
    Salary    int
    ManagerID int
}

var dilbert Employee

func main() {
    var otroTexto = " Miriam"
    fmt.Println( texto + otroTexto )
    
    fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // "Pointy-haired boss"

	id := dilbert.ID
	EmployeeByID(id).Salary = 0 // fired for... no real reason
}



func EmployeeByID(id int) *Employee { 
    fmt.Println(id)
    return &dilbert
}
