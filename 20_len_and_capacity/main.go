package main

import "fmt"

func main(){

	scores:= make([]int,0,5)

	fmt.Println(scores, len(scores), cap(scores))

	scores = append(scores, 100)
	fmt.Println("after appending  100", scores)

	scores = append(scores, 200, 3000)
	fmt.Println("after appending  200, 3000", scores)

	scores = append(scores, 45, 55)
	fmt.Println("after appending  45, 55", scores)

	//if incase we r exceeding the capacity, go grows the backnfg array(usyally doubles)
	scores = append (scores,60)
	fmt.Println("after appending 60", scores, len(scores), cap(scores))

	todos := []string{"do youtube", "workout"}

	more := []string{"learn golang"}

	//...

	todos = append(todos, more...)
	fmt.Println(todos)
}