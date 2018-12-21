package main

import (
	"fmt"
)

func main()  {
	var point = 9
	if point == 10 {
		fmt.Println("Grade A")
	} else if point > 5 {
		fmt.Println("Grade B")
	} else {
		fmt.Println("Grade C")
	}
}