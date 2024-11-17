//go:build debug
// +build debug

package main

import "fmt"

const debug = true

func main() {
	fmt.Println("Using debug-enabled version")
	if debug {
		fmt.Println("Debugging enabled")
	}
}
