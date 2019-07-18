package main

import "fmt"

// Prints out the index and the character of runes.
func main() {
	for i := 90; i < 98; i++ {
		fmt.Printf("index: %d char: %s\n", i, string(rune(i)))
	}
}
