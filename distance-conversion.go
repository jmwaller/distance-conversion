package main

import (
	"fmt"

	"github.com/mymath"
)

func main() {
	const MilesToKilometers = 1.60934
	const KilometersToMiles = 0.621371

	var kilos float64
	var miles float64

	fmt.Println("Please enter a number in kilometers.  I will store your value in memory at", &kilos, "identified by the variable kilos")
	n, err := fmt.Scanf("%f", &kilos)

	if n != 1 {
		fmt.Println(err)
	} else {
		fmt.Println("Computing miles from kilos:  RoundPlus((", kilos, "*", KilometersToMiles, "),2)")
		fmt.Println("Storing result in memory at", &miles, "identified by the variable miles")
		miles := mymath.RoundPlus((kilos * KilometersToMiles), 2)

		fmt.Println(kilos, "km is", miles, "miles!")
	}
	n = 0
	err = nil

	fmt.Println("Please enter a number in miles.  Re-using the miles variable in memory at", &miles, "to store the value you enter")
	n, err = fmt.Scanf("%f", &miles)

	if n != 1 {
		fmt.Println(err)
	} else {
		fmt.Println("Computing kilos from miles:  RoundPlus((", miles, "*", MilesToKilometers, "),2)")
		fmt.Println("Over-writing the kilos variable in memory at", &kilos)
		kilos = mymath.RoundPlus((miles * MilesToKilometers), 2)

		fmt.Println(miles, "miles is", kilos, "kilometers!")
	}
	fmt.Println("=======================================")

	fmt.Println("The memory address of the miles variable is", &miles)
	fmt.Println("The value of the miles variable is", miles)
	fmt.Println("Rounding miles value in place:  RoundInPlace(", &miles, ")")

	mymath.RoundInPlace(&miles)

	fmt.Println("The value of the miles variable is", miles)
	fmt.Println("The memory address of the miles variable is", &miles)

}
