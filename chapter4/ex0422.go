package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	words := []string{"山", "微笑み", "人類学者", "モグラの穴", "mountain", "イカとタコの足", "antholopologist", "タコの足は8本でイカの足は10本"}

	for _, word := range words {
		switch rc := utf8.RuneCountInString(word); rc {
		case 1, 2, 3, 4:
			fmt.Printf("「%s」の文字数は%dで、短い単語だ\n", word, rc)
		case 5:
			bc := len(word)
			fmt.Printf("「%s」の文字数は%dで、これはちょうど良い長さだ。ちなみにバイト数は%dだ。\n", word, rc, bc)
		case 6, 7, 8, 9:
		default:
			fmt.Printf("「%s」の文字数は%dで、とても長い!\n", word, rc)
		}
	}
}
