package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5 // setting a constant
	var investmentAmount, years float64 // multiple variable in same line (with no assignment)
	ExpectedReturnRate  := 5.5 // alternate assignment, automatically inferred type
	var textVar = `
hello,
this app calculates the future value of you investment.
Thanks
-------------------
	`
	fmt.Print(textVar)

	// fmt.Print("Enter investment amount, then years: ")
	// fmt.Scan(&investmentAmount, &years) // getting multiple value from the console input in single code line
	// fmt.Print("Enter rate of return: ")
	// fmt.Scan(&ExpectedReturnRate) // & signifies pointer the variable

	futureValue := investmentAmount * math.Pow((1+ExpectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow( (1+inflationRate/100) , years)
	tempVar := fmt.Sprintf("\n futureRealValue: %.2f \n", futureRealValue) // format and keep output in a variable

	// fmt.Println("futureRealValue: ", futureRealValue)
	fmt.Printf("futureValue: %v \n", futureValue)  // formatted output
	// fmt.Printf("futureRealValue: %.2f \n", futureRealValue) // control precision of the decimal output]
	fmt.Print(tempVar)
	outputText("Bye") // using you user defined functions
	}	


func outputText(text string){
	fmt.Println(text)
}