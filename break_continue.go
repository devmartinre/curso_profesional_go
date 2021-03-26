package main

import "fmt"

func main() {
	// break - continue
	for i := 0; i <= 10; i++ {
		if i == 5 {
			continue
		}

		if i == 8 {
			break
		}

		fmt.Println(i)
	}
}
