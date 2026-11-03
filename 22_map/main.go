package main
import "fmt"
func main (){
	//map [kayType]valueType
	ages := map[string]int{
		"sangam" : 65,
		"john" : 35,
	}

	fmt.Println(ages["sangam"], len(ages))

	var scores map[string]int

	fmt.Println(scores, scores["a"])

	scores = make(map[string]int)

	scores["math"] = 90
	fmt.Println(scores, scores["math"])

	users := map[string]string{
		"u1": "sangam",
		"u2": "Temi",
		"u3": "Raul",
	}

	fmt.Println(users)

	delete(users, "u3")
	delete(users, "u100") //no error

	fmt.Println(users)
}