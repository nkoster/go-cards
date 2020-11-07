package main

func main() {
	cards := readDeck("test.deck")
	// cards.print()
	cards.shuffle2(2)
	cards.print()
	// hand, remainingDeck := deal(cards, 5)
	// cards = remainingDeck
	// hand.print()
	// cards.print()
	// if err := cards.saveDeck("test.deck"); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// } else {
	// 	fmt.Println("deck saved")
	// }
	// newCards := cards.readDeck("test.deck")
	// newCards.print()
}
