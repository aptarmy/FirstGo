package main

import (
	"fmt"
	"github.com/aptarmy/FirstGo/newmath"
	"strconv"
)

func main() {
	var inputStr string
	var inputfloat float64

	fmt.Println("\n")
	fmt.Println("############## START #############")
	fmt.Print("Please Enter number to Cal Squeroot >> ")
	fmt.Scanln(&inputStr)
	inputfloat, _ = strconv.ParseFloat(inputStr, 64)
	fmt.Println("Result Squeroot : ", newmath.Sqrt(inputfloat))

	fmt.Print("Please Enter number to Cal Power >> ")
	fmt.Scanln(&inputStr)
	inputfloat, _ = strconv.ParseFloat(inputStr, 64)
	fmt.Println("Result Power : ", newmath.Power(inputfloat))
	fmt.Println("############### END ##############")
}
