package main

import "fmt"

func main() {

	fmt.Print("masukkan angka : ")

	var input int64
	var output string
	money := []int64{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	fmt.Scanf("%d", &input)
	temp := input
	cek := input % 100
	if cek > 0 && cek < 100 {
		temp = input - cek + 100
	}
	for i := 0; i < len(money); i++ {
		if temp/int64(money[i]) != 0 {
			output += fmt.Sprintf("RP.%v : %v \n", money[i], temp/int64(money[i]))
			temp %= money[i]
		}
	}
	fmt.Println("output: \n", output)

	fmt.Scanf("press enter to close")
	fmt.Scanf("%d", &input)

}
