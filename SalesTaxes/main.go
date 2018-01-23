package main

import "fmt"

func main() {

	var receipts []receipt
	fmt.Println("Inputs:")
	fmt.Println("+++++++++++++++++++++++++++++")
	for index := 0; index < noOfInputs; index++ {
		receipts = append(receipts, createReceipt())
		printReceipt(receipts[index])
		receipts[index].calculateReceipt()

	}
	fmt.Println("Sales Report:")
	fmt.Println("Outputs")
	fmt.Println("+++++++++++++++++++++++++++++")
	for index := 0; index < noOfInputs; index++ {
		printReceipt(receipts[index])
	}

}
