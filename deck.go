package main

import "fmt"

// CREATE A NEW TYPE OF DECK WHICH IS A SLICE OF STRINGS
type deck []string

//(d deck) = Reciver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
