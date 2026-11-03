package main 
import "fmt"

func add(a int, b int) int{
	return a + b
}


func SumAndProduct(a int, b int) (int, int){
	sum := a + b
	product := a * b

	return sum, product
}
func main (){
	sum1 := add(10, 20)
	s, p := SumAndProduct(6, 5)
	fmt.Println(sum1, s, p)

	onlySum, _ := SumAndProduct(10, 2)
		fmt.Println(onlySum)

		_ , onlyProduct := SumAndProduct(20, 2)
		fmt.Println(onlyProduct)
	
}