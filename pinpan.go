package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {

		pin := i%3 == 0
		pan := i%5 == 0

		switch {
		case pin && pan:
			fmt.Println("Pin Pan")
		case pin && !pan:
			fmt.Println("Pin")
		case !pin && pan:
			fmt.Println("Pan")
		}
	}
}
