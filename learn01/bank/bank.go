package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const balanceFileName = "balance.txt"



func main(){
	fmt.Println(randomdata.SillyName())
	accountBalance, err := fileops.GetFloatFromFile(balanceFileName)
	if err != nil{
		fmt.Println(err)
		// panic(err)
	}
	for {
		
		present_options()
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		fmt.Printf("\nYour choice: %v\n", choice)

		
		
		if (choice == 1) {
			fmt.Printf("Your balance is: %.2f\n", accountBalance)
		}else if (choice ==2){
			depositAmount:= 0.0
			fmt.Print("How much do you want to deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0{
				fmt.Println("Amount entered was less that equal to zero")
				continue
			}
			accountBalance+= depositAmount
			fileops.WriteFloatToFile(balanceFileName, accountBalance)
			fmt.Printf("\nYour Balance: %v\n", accountBalance)
		} else if choice == 3 {
			withdrawAmount:= 0.0
			fmt.Print("How much do you want to deposit: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0{
				fmt.Println("Amount entered was less that equal to zero")
				continue
			}
			if withdrawAmount >= accountBalance{
				fmt.Println("Amount entered was more than account balance")
				continue
			}
			accountBalance-= withdrawAmount
			fileops.WriteFloatToFile(balanceFileName, accountBalance)
			fmt.Printf("\nYour Balance: %v\n", accountBalance)
		} else {
			fmt.Println("Good Bye !")
			break
		}
	}
	
	// choice := 1
	// switch choice{
	// case 1:
	// 	fmt.Printf("Your balance is: %.2f\n", accountBalance)
	// 	fmt.Printf("Your balance is: %.2f\n", accountBalance)
	// case 2:
	// 	fmt.Println("2 was selected")
	// default:
	// 	fmt.Println("Good Bye !")
	// 	return

	// }

}
