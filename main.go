package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.print()
	//cards.saveToFile("my_cards")
	//fmt.Println(cards.toString())
	/*hand, remaingingCards := deal(cards, 5)
	hand.print()
	remaingingCards.print()*/
}
