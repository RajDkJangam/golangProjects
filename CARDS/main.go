package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	//card := newCard() // : only for intilization
	//cards := deck{"Ace of Diamonds", newCard()} //Slice
	//cards = append(cards, "Six of Spades")
	cards := newDeck()
	cards.saveToFile("my_cards")

	hand, _ := deal(cards, 5)
	hand.print()
	fmt.Println(cards.toString())

	cardsFile := newDeckFromFile("my_cards")
	cardsFile.shuffle()
	cardsFile.print()
	//cards.print()
	//remainingCards.print()
}
