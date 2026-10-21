package main
import "fmt"
func add (x int, y int) int{
  return x + y
}

func addd (x , y int) int{
	return x + y
}

func main (){
  fmt.Println(addd(10, 20))
}