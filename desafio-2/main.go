package main

import "fmt"

func main() {
	for i := range 100 {
		num := i + 1

		if num%3 == 0 {
			fmt.Println("Pin")
		}

		if num%5 == 0 {
			fmt.Println("Pan")
		}
	}
}
