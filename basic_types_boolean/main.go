package main

import (
	"fmt"
)

func main(){
	isLogged := true;
	isAdmin := false;
	hasSuscribed := false;

	canOpenDashboard := isLogged && hasSuscribed

	canDeletePost := isAdmin || (isLogged && hasSuscribed)

	fmt.Println(canOpenDashboard, canDeletePost)

	age := 20
	isAdult := age > 18
	fmt.Println(isAdult)
}