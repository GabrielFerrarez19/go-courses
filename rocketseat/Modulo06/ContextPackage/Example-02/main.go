package main

import "fmt"

type (
	MyIntWrapper      int
	MyOtherIntWrapper int
)

func main() {
	m := make(map[any]any)
	m[1337] = 3000
	fmt.Println(m[1337])
	m[1337] = 2000
	fmt.Println(m[1337])

	m[MyIntWrapper(1337)] = 10
	m[MyOtherIntWrapper(1337)] = 200

	fmt.Println(m[1337], m[MyOtherIntWrapper(1337)], m[MyIntWrapper(1337)])
}
