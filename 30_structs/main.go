
package main

import "fmt"

// A struct groups related fields into one type
type User struct {
    ID    int
    Name  string
    Email string
    Age   int
}

func main() {
    u1 := User{
        ID:    1,
        Name:  "Sangam",
        Email: "sulaimantemiloluwa24@gmail.com",
        Age:   24,
    }

    fmt.Println(u1)
    fmt.Println(u1.ID, u1.Email)

    // Modify the user's age
    u1.Age = 200

    fmt.Println("Updated age:", u1.Age)

	u2 := User{
		Name: "Temzy",
		Email: "jsjasjkfadf",
	}

	fmt.Println("partial user:", u2)
}
