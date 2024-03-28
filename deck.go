package main

import (
	"fmt"
	"math/rand"
)

type deck []string

type Card struct {
	Name string
}

func newDeck() deck{
	cards := deck{}
	cardSuits := []string{"♠️","♥️","♦️","♣️"}
	cardValues := []string{"A","2","3","4","5","6","7","8","9","10","J","Q","K"}

	for _, suit :=range cardSuits{
		for _, value := range cardValues{
			cards = append(cards, suit + value)
		}
	}
	return cards
}

//show the list
func (d deck) ShowList() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//shuffle and get 1
func (d deck) Deal(size int) deck{
	// randId := rand.Intn(len(d))
	// randCard := d[randId]
	// fmt.Println("gia tri sau khi shuffle ",randCard)
	
	if size > len(d){
		fmt.Println("not enough card in the deck")
		return d
	}
	newHand := newDeck()
	for i:= 0, i<size, i++{
		
	}

	return d[:size]
}


func (d deck) shuffle() {
	for i := range d {
        j := rand.Intn(i + 1) // Random index between 0 and i (inclusive)
        d[i], d[j] = d[j], d[i]
    }
}



// type Deck struct {
// 	cards []Card
// }

// func NewDeck() Deck {
// 	cards := []string{"AK", Card}
// 	// cardsk := make([]Card, 0)
// 	return Deck{cards: cards}
// }