package main

import "fmt"

const x int64 = 10

const (
// idKey   = "id"
// nameKey = "name"
)

const z = 20 * 10

func main() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// x = x + 1 // => produces an error that you can't reassign x
	//y = "bye" // => produces an error that you can't reassign y

	fmt.Println(x)
	fmt.Println(y)
}
