package main

import (
	"errors"
	"fmt"
	"os"
)


func main() {
	fmt.Println("Profit Calculator")
	// var revenue, expense, taxRate float64

	// fmt.Print("revenue: ")
	// fmt.Scan(&revenue)
	revenue, err := take_user_input("revenue")
	if err != nil{
		panic(err)
	}
	// fmt.Print("expense: ")
	// fmt.Scan(&expense)
	expense, err := take_user_input("expense")
	if err != nil{
		panic(err)
	}
	// fmt.Print("taxRate: ")
	// fmt.Scan(&taxRate)
	taxRate, err := take_user_input("taxRate")
	if err != nil{
		panic(err)
	}

	earningBeforeTax := revenue - expense
	// profit := earningBeforeTax - (earningBeforeTax * (taxRate/100))
	// ratio := earningBeforeTax/profit
	profit, ratio:= calculate_profit(earningBeforeTax, taxRate)
	writeDataToFile(profit, ratio)
	fmt.Println(profit)	
	fmt.Println(ratio)
}

func take_user_input(prompt string) (inputVal float64, err error){
	fmt.Printf("%v: ", prompt)
	fmt.Scan(&inputVal)
	if inputVal <= 0{
		return inputVal, errors.New("Input value was less that or equal to zero")
	}
	return inputVal, nil
}

func calculate_profit(earningBeforeTax float64, taxRate float64) (profit float64, ratio float64){
	profit = earningBeforeTax - (earningBeforeTax * (taxRate/100))
	ratio = earningBeforeTax/profit
	return profit, ratio
}

func writeDataToFile(profit float64, ratio float64) (error){
	// profitText := strconv.FormatFloat(profit, 'b', 2, 64)
	// ratioText := strconv.FormatFloat(ratio, 'b', 2, 64)

	textToWrite := fmt.Sprintf("profit: %.2f, ratio: %.3f", profit)
	os.WriteFile("test.txt", []byte(textToWrite), 0644)
	return nil
}
