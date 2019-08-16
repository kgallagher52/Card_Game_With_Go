package main

import "fmt"

// CREATE A NEW TYPE OF DECK WHICH IS A SLICE OF STRINGS
type deck []string

//(d deck) = Reciver
//deck: any variable that has access to deck can use print
//d: is the reference to deck type variable type deck []string
//print(): the function name
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
