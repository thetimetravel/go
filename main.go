package qg

import "fmt"

type QG struct {
	Name string
	Age int
} 
func init()  {
	qg:=QG{
		Name: "wwwwwww",
		Age:  444,
	}
	qg.Show_info()
	fmt.Println("good:github")
}

func (qg *QG)Show_info()  {
	fmt.Println("good")
	fmt.Println("info: name",qg.Name," ,age:  ",qg.Age)
}
