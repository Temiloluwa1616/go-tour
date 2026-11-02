package main
import "fmt"
func main(){
	views := []int{10, 20, 45, 50, 60}

	//for range
	total := 0

	for i, v:= range views {
		fmt.Println("day", i, "views", v)
		total = total + v
	}

	fmt.Println(total)
}