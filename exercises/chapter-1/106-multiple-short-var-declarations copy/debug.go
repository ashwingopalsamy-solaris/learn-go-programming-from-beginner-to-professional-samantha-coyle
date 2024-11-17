package main

import "fmt"

// Default value for debug constant, will be overwritten at compile time.
const debug = false

func main() {
    fmt.Println("Using debug version:", debug)
    if debug {
        fmt.Println("Debugging enabled")
    } else {
        fmt.Println("Debugging disabled")
    }
}
