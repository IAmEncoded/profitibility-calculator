package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64


	fmt.Print("Please enter your yearly revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Please enter your yearly expenses: ")
	fmt.Scan(&expenses)
	
	fmt.Print("Please enter your yearly tax rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	earningsAfterTax := float64(earningsBeforeTax) * (1 - taxRate / 100)
	ratio := earningsBeforeTax / earningsAfterTax

	fmt.Print("EBT: ")
	fmt.Println(earningsBeforeTax)

	fmt.Print("Earnings After Tax: ")
	fmt.Println(earningsAfterTax)

	fmt.Print("Ratio of earnings to profit: ")
	fmt.Println(ratio)

}