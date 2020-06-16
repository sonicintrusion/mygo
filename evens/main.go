package main

import (
	"fmt"
)

func main() {
	var oddeven string
	var values []int

	for i := 0; i <= 10; i++ {
		values = append(values, i)
	}

	for _, value := range values {
		if value%2 == 0 {
			oddeven = "even"
		}
		if value%2 != 0 {
			oddeven = "odd"
		}
		fmt.Println(value, "is", oddeven)
	}

}
