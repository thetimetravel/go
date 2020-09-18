package qg

import "fmt"

type QG struct {
	name string
	age int
} 
//func init()  {
//	qg:=QG{
//		name: "xq",
//		age:  33,
//	}
//	qg.Show_info()
//	fmt.Println("good:github")
//}

func (qg *QG)Show_info()  {
	fmt.Println("info: name",qg.name," ,age:  ",qg.age)
}
