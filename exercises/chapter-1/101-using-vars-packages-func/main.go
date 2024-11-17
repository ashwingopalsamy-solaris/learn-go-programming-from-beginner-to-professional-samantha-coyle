package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	r := rnd.Intn(5) + 1
	stars := strings.Repeat("*", r)
	fmt.Println(stars)
}
