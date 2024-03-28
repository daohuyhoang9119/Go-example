package main

import "fmt"
func main(){
	// cards := deck{"AK", newCard()}
	cards := newDeck()
	// cards = append(cards, "QQ")
	// cards = append(cards, "KK")
	// cards = append(cards, "99")
	// cards = append(cards, "88")
	// cards = append(cards, "67")
	fmt.Println("show list")
	cards.ShowList()
	fmt.Println("shuffle and get 1")
	cards.ShuffleAndGet1()
	fmt.Println("shuffle all the card inside the deck and show list")
	cards.shuffle()
	cards.ShowList()
}

func newCard() string{
	return "JJ"
}