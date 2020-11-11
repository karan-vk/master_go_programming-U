package main

import "fmt"

func main() {
	const (
		c11 = iota
		c22 
		c33
	)
	fmt.Println(c11,c22,c33)
}
