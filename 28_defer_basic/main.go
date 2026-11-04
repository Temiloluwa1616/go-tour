
package main

import (
    "fmt"
    "errors"
)

func main() {
    fmt.Println("Case 1: success")

    if err := doWork(true); err != nil {
        fmt.Println("error:", err)
    }

    fmt.Println("Case 2: fail early")

    if err := doWork(false); err != nil {
        fmt.Println("error:", err)
    }
}

func doWork(success bool) error {
    fmt.Println("start: resource acquired")

    defer fmt.Println("cleanup: resource released")

    if !success {
        return errors.New("something went wrong i am returning early")
    }

    fmt.Println("work: doing s omething important")
    fmt.Println("work: this work is done")

    return nil
}
