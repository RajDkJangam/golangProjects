package main

import (
	"testing"
)

func TestNewReceipt(t *testing.T) { // t is Test Handler
	r := createReceipt()

	if len(r.items) != 5 {
		t.Errorf("Expected count of items of 5 in Reciept, but got %v", len(r.items))
	}
}

func TestOutputCount(t *testing.T) { // t is Test Handler
	var receipts []receipt
	for index := 0; index < noOfInputs; index++ {
		receipts = append(receipts, createReceipt())
		receipts[index].calculateReceipt()
	}

	if len(receipts) != 3 {
		t.Errorf("Expected count of outputs is 3, but got %v", len(receipts))
	}
}

func TestTaxForNonTaxable(t *testing.T) { // t is Test Handler
	testItem := item{
		itemType:   "book",
		shelfPrice: 10.00,
		isImported: false,
	}
	taxamount := testItem.calculateTax()

	if taxamount != 0 {
		t.Errorf("Expected Tax should be zero for non taxable item , but got %v", taxamount)
	}
}

func TestTaxForTaxable(t *testing.T) { // t is Test Handler
	testItem := item{
		itemType:   "bottle of perfume",
		shelfPrice: 19.99,
		isImported: false,
	}
	taxamount := testItem.calculateTax()

	if taxamount == 0 {
		t.Errorf("Expected Tax should not be zero for taxable item , but got %v", taxamount)
	}
}

func TestImportDutyForTaxable(t *testing.T) { // t is Test Handler
	testItem := item{
		itemType:   "bottle of perfume",
		shelfPrice: 30.00,
		isImported: true,
	}
	taxamount := testItem.calculateTax()

	if taxamount < 0.15*(testItem.shelfPrice) {
		t.Errorf("Expected Import Duty and Tax should not less than 15 percentage for taxable item , but got %v", taxamount)
	}
}

func TestImportDutyForNonTaxable(t *testing.T) { // t is Test Handler
	testItem := item{
		itemType:   "bottle of perfume",
		shelfPrice: 30.00,
		isImported: true,
	}
	taxamount := testItem.calculateTax()

	if taxamount < 0.05*(testItem.shelfPrice) {
		t.Errorf("Expected Import Duty and Tax should not less than 5 percentage for non taxable item , but got %v", taxamount)
	}
}
