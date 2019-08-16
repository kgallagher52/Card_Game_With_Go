package main

import "fmt"

// CREATE A NEW TYPE OF DECK WHICH IS A SLICE OF STRINGS
type deck []string //1

//Create and return list of cards
func newDeck() deck {
	cards := deck{}

	//Slice of type string
	cardSuites := []string{"Hearts", "Spades", "Dimonds", "Clubs"}
	cardValues := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}

	//Loop through both suites and values and append value and suites
	for _, suits := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suits)
		}
	}
	return cards
}

//(d deck) = Reciver
//deck: any variable that has access to deck can use print
//d: is the reference to deck type variable type deck []string
//print(): the function name
// d: refernce of cards.
// deck: is the type
func (d deck) print() { //4
	for i, card := range d {
		fmt.Println(i, card)
	}
}
