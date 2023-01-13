package main

import "fmt"

func main() {
	fmt.Println("1 + 1 = ", 1.0+1.0)

	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[0])
	fmt.Println("Hello " + "World!")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)
	fmt.Println(32132 * 42452)
}
