
package main

import "fmt"

type User struct {
    Name string
    Age  int
}

func main() {
    u := User{Name: "Temi", Age: 20}

    fmt.Println(u.intro())
}

// Value receiver
func (u User) intro() string {
    return fmt.Sprintf("Hi, I am %s", u.Name)
}
