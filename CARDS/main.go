package main

func main() {
	cards := newDeck()

	cards.saveToFile("my_cards")

	// fmt.Println(cards.toString())
}
