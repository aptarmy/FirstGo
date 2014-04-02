package main

import (
	"fmt"
	"github.com/aptarmy/FirstGo/newmath"
	"strconv"
)

func main() {
	var inputStr string
	var inputfloat float64

	fmt.Print("Please Enter number to Cal Squeroot >> ")
	fmt.Scanln(&inputStr)

	inputfloat, _ = strconv.ParseFloat(inputStr, 64)
	fmt.Println("Result : ", newmath.Sqrt(inputfloat))
}
