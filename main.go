package main

import (
	"fmt"
	"strings"
)

//Enter the temperature
//from unit
//to the unit you want to convert
//display the result

func main() {
	var temp float64
	var fromUnit, toUnit string

	fmt.Print("Enter the Temperature: ")
	fmt.Scanln(&temp)
	fmt.Print("From Unit (C, F, K) : ")
	fmt.Scanln(&fromUnit)
	fromUnit = strings.ToUpper(fromUnit)
	fmt.Print("To Unit (C, F, K) : ")
	fmt.Scanln(&toUnit)
	toUnit = strings.ToUpper(toUnit)

	convertedTemp := 0.0

	switch fromUnit{
	case "C" :
		if toUnit == "F"{
			convertedTemp = temp*9/5 + 32
		}else if toUnit == "K"{
			convertedTemp = temp + 273.15
		}else{
			fmt.Println("Invalid Input.")
		}

	case "F" :
		if toUnit == "C"{
			convertedTemp = (temp-32) *5/9
		}else if toUnit == "K"{
			convertedTemp = (temp-32) *5/9 +273.15
		} else{
			fmt.Println("Invalid Input.")
		}

	case "K":
		if toUnit == "C"{
			convertedTemp = temp-273.15
		}else if toUnit == "F"{
			convertedTemp = (temp- 273.15)*9/5 +32
		}else{
			fmt.Println("Invalid Input.")
		}

	default:
		fmt.Println("Invalid Input(fromUnit).")
	}

	fmt.Printf("%v %s is %.2f %s\n", temp, fromUnit, convertedTemp, toUnit)

}
