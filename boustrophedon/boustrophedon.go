package main

import (
	"strings"
	"fmt"
)

var testText = `A modern example of boustrophedonics is the numbering scheme of sections within survey townships
in the United States and Canada. In both countries, survey townships are divided into a 6-by-6 grid
of 36 sections. In the U.S. Public Land Survey System, Section 1 of a township is in the northeast
corner, and the numbering proceeds boustrophedonically until Section 36 is reached in the southeast
corner.[8] Canada's Dominion Land Survey also uses boustrophedonic numbering, but starts at the southeast
corner.[9]
Source: https://en.wikipedia.org/wiki/Boustrophedon
`

var lookup = make(map[rune]rune)

func init() {
	o := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	m := []rune{'ɒ', 'd', 'ɔ', 'b', 'ɘ', 'ʇ', 'ǫ', 'ʜ', 'i', 'Ⴑ', 'ʞ', 'l', 'm', 'n', 'o', 'q', 'p', 'ɿ', 'ƨ', 'ƚ', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'ᙠ', 'Ɔ', 'ᗡ', 'Ǝ', 'ᖷ', 'Ꭾ', 'H', 'I', 'Ⴑ', 'ᐴ', '⅃', 'M', 'И', 'O', 'ꟼ', 'Ọ', 'Я', 'Ƨ', 'T', 'U', 'V', 'W', 'X', 'Y', 'Ƹ'}
	for i, l := range o {
		lookup[l] = m[i]
	}
}

func main() {
	fmt.Println(testText)
	fmt.Println(Boustrophedon(testText, false))
	fmt.Println(Boustrophedon(testText, true))
}

func Boustrophedon(s string, mirrorLetters bool) string {
	apply := Reverse
	if mirrorLetters {
		apply = ReverseMirrored
	}

	lines := strings.Split(s, "\n")
	for i := range lines {
		if i%2 == 1 {
			lines[i] = apply(lines[i])
		}
	}
	return strings.Join(lines, "\n")
}

func Reverse(s string) string {
	var reverse string
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}

func ReverseMirrored(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if k, ok := lookup[runes[i]]; ok {
			runes[i] = k
		}
		if k, ok := lookup[runes[j]]; ok {
			runes[j] = k
		}
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
