package main

import (
	"fmt"
	"unicode/utf8"
)

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func printCharsAndBytes(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func length(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}

func main() {
	h := "hello"
	fmt.Println(mutate([]rune(h)))

	/*
		word1 := "Señor"
		length(word1)
		word2 := "Pets"
		length(word2)
	*/

	/*
		runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
		str := string(runeSlice)
		fmt.Println(str)
	*/

	/*
		byteSlice := []byte{67, 97, 102, 195, 169}
		str := string(byteSlice)
		fmt.Println(str)
	*/

	/*
		byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
		str := string(byteSlice)
		fmt.Println(str)
	*/

	/*
		name := "Hello world"
		printBytes(name)
		fmt.Println()
		printChars(name)
		fmt.Println()

		name = "Señor"
		printBytes(name)
		fmt.Println()
		printChars(name)
		fmt.Println()

		printCharsAndBytes(name)
		fmt.Println()
	*/
}
