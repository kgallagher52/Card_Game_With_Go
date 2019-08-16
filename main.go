//First lesson
// package main

// // FMT is for formating, commonly used with Println
// import "fmt"

// func main() {
// 	// var card string = "Ace of Spades"
// 	// Short hand for line above
// 	card := newCard()
// 	// := is only used when creating a new variable
// 	fmt.Println(card)
// }

// func newCard() string { // Function decleration
// 	return "5 of diamonds"
// }

package main

import "fmt"

func main() {
	//Slice of type string
	cards := []string{"Ace of diamonds", newCard()}
	cards = append(cards, "six of spades")
	fmt.Println(my_cards)
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// func newCard() string { // Function decleration
// 	return my_cards
// }
