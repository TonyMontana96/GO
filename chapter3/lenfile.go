package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("variables.exe")
	if err != nil {
		log.Println(err.Error())
		return
	}
	info, err := f.Stat()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(info.Size())
}
