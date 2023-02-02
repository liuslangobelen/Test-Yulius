package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("no 2")
	var input1 string
	var input2 string
	fmt.Print("GIVEN INPUT 1 --> ")
	fmt.Scanln(&input1)
	fmt.Print("GIVEN INPUT 2 --> ")
	fmt.Scanln(&input2)
	input1 = strings.ToLower(input1)
	input2 = strings.ToLower(input2)
	result := false
	lenMax := len(input1)
	max := input1
	lenMin := len(input1)
	min := input1
	if len(input1) > len(input2) {
		lenMin = len(input2)
		min = input2
	} else {
		lenMax = len(input2)
		max = input2
	}
	diffCount := 0
	if lenMax-lenMin <= 1 {
		loop1 := -1

		for i := 0; i < lenMax; i++ {
			loop1++
			if min[loop1] != max[i] {
				if lenMax != lenMin {
					loop1--
				}
				diffCount++
			}
			if diffCount > 1 {
				break
			}
		}
	}
	if diffCount == 1 {
		result = true
	}

	fmt.Println("RESULT : ", result)

	fmt.Scanf("press enter to close")
	fmt.Scanf("%d", &input1)

}
