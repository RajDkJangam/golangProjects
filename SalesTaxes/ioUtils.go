package main

import (
	"fmt"
)

//Item's Input Print Function
func iPrint(i item) {
	if i.isImported {
		fmt.Println("1 imported", i.itemType, "at", i.shelfPrice)
	} else {
		fmt.Println("1", i.itemType, "at", i.shelfPrice)
	}
}

//Item's Output Print Function
func oPrint(i item) {
	if i.isImported {
		fmt.Println("1 imported", i.itemType, ":", i.shelfPrice)
	} else {
		fmt.Println("1", i.itemType, ":", i.shelfPrice)
	}
}

//Printing of the receipt
func printReceipt(r receipt) {
	if r.totalAmount == 0 {
		fmt.Println("Input", r.receiptNo, ":")
		for _, it := range r.items {
			iPrint(it)
		}
		fmt.Println()
	} else {
		fmt.Println("Output", r.receiptNo, ":")
		for _, it := range r.items {
			oPrint(it)
		}
		fmt.Printf("Sales Taxes:%0.2f\n", r.totalTax)
		fmt.Printf("Total Amount:%0.2f\n", r.totalAmount)
		fmt.Println()
	}
}
