package main

import "fmt"

func divide (a int, b int) (john int, sangam int){
	john = a/b
	sangam = a + b
	return
}


func main (){

	q, r := divide (10,10)
	fmt.Println(q, r)
}