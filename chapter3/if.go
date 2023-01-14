package main

import "fmt"

func main() {
	for b := 34; b <= 100; b++ {
		if b%2 == 0 {
			fmt.Println(b, "true")
		} else {
			fmt.Println(b, "false")
		}
	}
}
