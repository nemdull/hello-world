package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var data []int
	x := []int{}
	fmt.Println(data == nil) //true
	fmt.Println(x == nil)    //false
	calc()
	calc2()
	rune_ex()
}

func calc() {
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2]
	z := x[2:]
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, "i", "j", "k")
	x = append(x, "y")
	z = append(z, "y")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

}
func calc2() {
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d") // len=4 cap=5

	y := x[:2]   // サブスライス:len=2 cap=5（共有・危険）
	z := x[:2:2] // フルスライス:len=2 cap=2（cap制限・安全）

	y = append(y, "Y", "Y") // 共有配列に追記→xが汚染され得る
	z = append(z, "Z")      // cap不足→新配列に退避→xは無傷

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

func rune_ex() {
	s := "こんにちは"                           // 1文字=3バイト、合計15バイト
	fmt.Println(len(s))                    // 15（バイト数）
	fmt.Println(utf8.RuneCountInString(s)) // 5（rune数）

	// インデックスはバイト単位
	fmt.Printf("%x\n", s[0]) // e3（先頭バイト）

	// range はrune単位
	for i, r := range s {
		fmt.Printf("i=%d r=%c\n", i, r)
	}

	// 文字列⇄runeスライス
	rs := []rune(s)
	rs[0] = 'さ'
	fmt.Println(string(rs)) // 「さんにちは」

	for i, r := range rs {
		fmt.Printf("i=%d r=%c\n", i, r)
	}
}
