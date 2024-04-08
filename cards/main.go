// // savinng data to the hard drive (run file in terminal before pushing)
// // reading from the hard drive
package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
}

// error handling
// package main

// func main() {
// 	// cards := newDeck()
// 	// cards.saveToFile("my_cards")

// 	cards := newDeckFromFile("my_cards")
// 	cards.print()
// }