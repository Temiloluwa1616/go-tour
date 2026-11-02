package main

import (
	"fmt"
)

func main (){
	//fixed and cant grow 
	var marks [3]int

	marks[0] = 10
	marks[1] = 20
	marks[2] = 50

	fmt.Println(marks)

	//array literal
	res := [5]int{2,3,4,5,6}

	fmt.Println(len(res))
}