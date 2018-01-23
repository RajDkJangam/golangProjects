package main

type receipt struct {
	receiptNo   int
	items       []item
	totalTax    float64
	totalAmount float64
}

//Calculation of Receipt's Item Tax, Total Tax and Total Amount
func (r *receipt) calculateReceipt() {
	(*r).totalTax = 0
	(*r).totalAmount = 0
	for in, it := range (*r).items {
		itemTax := it.calculateTax()
		it.includeTax(itemTax)
		(*r).totalTax = addFunction((*r).totalTax, itemTax)
		(*r).totalAmount = addFunction((*r).totalAmount, it.shelfPrice)
		(*r).items[in] = it
	}
}

//Creation of Receipt
func createReceipt() receipt {
	var r receipt
	r.receiptNo = receiptCounter
	for in := 0; in < noOfItems; in++ {
		r.items = append(r.items, newItem())
	}
	for in := 0; in < noOfImportedItems; in++ {
		r.items = append(r.items, newImportedItem())
	}
	receiptCounter++
	return r
}
