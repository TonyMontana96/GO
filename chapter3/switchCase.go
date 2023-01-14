package main

import "log"

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
		switch i {
		case 0:
			log.Println("Zero")
		case 1:
			log.Println("One")
		case 2:
			log.Println("Two")
		case 3:
			log.Println("Three")
		case 4:
			log.Println("Four")
		case 5:
			log.Println("Five")
		case 6:
			log.Println("Six")
		case 7:
			log.Println("Seven")
		case 8:
			log.Println("Eight")
		case 9:
			log.Println("Nine")
		case 10:
			log.Println("Ten")
		}
	}
}
