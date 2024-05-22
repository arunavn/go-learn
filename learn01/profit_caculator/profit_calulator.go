package main

import "fmt"


func main() {
	fmt.Println("Profit Calculator")
	var revenue, expense, taxRate float64

	fmt.Print("revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("expense: ")
	fmt.Scan(&expense)

	fmt.Print("taxRate: ")
	fmt.Scan(&taxRate)

	earningBeforeTax := revenue - expense
	profit := earningBeforeTax - (earningBeforeTax * (taxRate/100))
	ratio := earningBeforeTax/profit
	fmt.Println(profit)	
	fmt.Println(ratio)
}
