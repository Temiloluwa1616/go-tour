package main

import (
	"fmt"
)

func main (){
	// common collection type
	// dynamic and can grow
	// []type{...}

	results := []string{"temiloluwa", "sulaiman"}
	fmt.Println(results, results[0], results[len(results)-1])

	results[1] = "priya"
	fmt.Println(results)


	var nums []int

	nums = append(nums, 10)
	nums = append(nums, 20,30)

	fmt.Println(nums)
}