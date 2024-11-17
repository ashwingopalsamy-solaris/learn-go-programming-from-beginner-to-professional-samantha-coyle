package main

import (
	"fmt"
	"time"
)

var (
	foo string = "ash"
	bar string = "win"
)

func main() {
	var t time.Time = time.Now()
	fmt.Println(t)
	fmt.Println(foo + bar)
}
