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
	fmt.Println("-------SHOW LIST --------")
	cards.ShowList()
	fmt.Println("------------DEAL------------")
	hand := cards.Deal(12)
	fmt.Println(hand)
	fmt.Println("shuffle all the card inside the deck and show list")
	cards.shuffle()
	cards.ShowList()
}

func newCard() string{
	return "JJ"
}