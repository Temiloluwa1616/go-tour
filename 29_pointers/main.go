package main
import "fmt"
func main(){
	// store the memory address of any val
	// &x -> adress of x (makes a pointer)
	// *p -> deref(go to that address and read/write)

	score := 10
	fmt.Println("before", score)

	addScore(&score)
	fmt.Println("after", score)
}

func addScore(score *int){

	*score = *score + 5
}