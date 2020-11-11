package main

import "fmt"

// import (
// 	"fmt"
// )

//Struct
type Person struct {
	name string
	age  int
}

func main() {
	var me Person
	me.age = 1
	// balances := map[string]float64{
	// 	"USD": 150.45,
	// }
	// fmt.Printf("%T \n",me)

	// Pointer
	ptr := &me
	println(ptr)
	f()

}
func f() {
	fmt.Println("Hello")
}
