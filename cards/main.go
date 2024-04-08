// // byte slices, deck to string, joining a slice of strings
package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())
}

// // savinng data to the hard drive (run file in terminal before pushing)
// // reading from the hard drive
// package main

// func main() {
// 	cards := newDeck()
// 	cards.saveToFile("my_cards")
// }

// error handling
// package main

// func main() {
// 	// cards := newDeck()
// 	// cards.saveToFile("my_cards")

// 	cards := newDeckFromFile("my_cards")
// 	cards.print()
// }