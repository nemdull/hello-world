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

	switch_ex()
	break_ex()
}

func switch_ex() {
	words := []string{"hi", "salutations", "hello"}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "は短い単語です")
		case wordLen > 10:
			fmt.Println(word, "は長すぎる単語です")
		default:
			fmt.Println(word, "はちょうど良い長さです")
		}
	}
}

func break_ex() {
loop:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, ":even")
		case i%3 == 0:
			fmt.Println(i, ":mod 3 but not mod 2")
		case i%7 == 0:
			fmt.Println(i, ":end of loop")
			break loop //label
		default:
			fmt.Println(i, ":boaring number")
		}
	}
}
