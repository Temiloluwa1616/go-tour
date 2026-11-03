package  main
import "fmt"


func main (){
	points := map[string]int{
		"a" : 10,
		"b" : 0,
	}

	fmt.Println("a", points["a"])
	fmt.Println("b", points["b"])
	fmt.Println("c", points["c"])

	valB, okB := points["b"]
	fmt.Println(valB, okB)

	valC, okC := points["c"]
	fmt.Println(valC, okC)

	if val, ok := points["a"]; ok {
		fmt.Println(val)
	} else{
		fmt.Println("c key is not present in the map")
	}

	prices := map[string]int{
		"xyz" : 500,
		"def" : 1800,
	}

	total := 0
	for item, price := range prices {
		fmt.Println(item,price)
		total = total + price
	}
	fmt.Println(total)

	for item := range prices{
		fmt.Println(item)
	}
}