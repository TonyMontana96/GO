package main

import "log"

var x = "Hello World"

func main() {
	log.Println(x)
}

func f() {
	log.Println(x)
}
