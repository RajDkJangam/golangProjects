package main

type item struct {
	itemType   string
	shelfPrice float64
	isImported bool
}

//Creation of new Item
func newItem() item {
	newPostion := randomGenerator()
	nI := item{
		itemType:   itemTypes[newPostion],
		shelfPrice: shelfPrices[newPostion],
		isImported: false,
	}
	return nI
}

//Creation of Imported Item
func newImportedItem() item {
	newPostion := randomGenerator()
	nII := item{
		itemType:   itemTypes[newPostion],
		shelfPrice: shelfPrices[newPostion],
		isImported: true,
	}
	return nII
}

//Calculation of Tax before rounding function
func (i item) taxBeforeRoundingFunction() float64 {
	var taxBeforeRounding float64
	switch i.itemType {
	case "book", "chocolate bar", "box of chocolates", "packet of headache pills":
		taxBeforeRounding = float64(0)
		break
	default:
		taxBeforeRounding = float64(10 * i.shelfPrice / 100)
	}

	if i.isImported {
		importDuty := 5 * i.shelfPrice / 100
		taxBeforeRounding = taxBeforeRounding + importDuty
	}

	return taxBeforeRounding

}

//Tax Calculation Function
func (i item) calculateTax() float64 {
	return roundingFunction(i.taxBeforeRoundingFunction())
}

//Updation of Shelf Price with Tax Amount
func (i *item) includeTax(taxAmount float64) {
	(*i).shelfPrice = addFunction((*i).shelfPrice, taxAmount)
}
