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

func main() {
	//Slice of type string
	cards := deck{"Ace of diamonds", newCard()}
	cards = append(cards, "six of spades")

	cards.print() // Getting this from the reciver in deck.go
}

func newCard() string { // Function decleration
	return "sssss"
}
