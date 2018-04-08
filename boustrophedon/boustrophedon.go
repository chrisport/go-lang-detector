package boustrophedon

import (
	"strings"
)


var lookup = make(map[rune]rune)

func init() {
	o := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	m := []rune{'ɒ', 'd', 'ɔ', 'b', 'ɘ', 'ʇ', 'ǫ', 'ʜ', 'i', 'Ⴑ', 'ʞ', 'l', 'm', 'n', 'o', 'q', 'p', 'ɿ', 'ƨ', 'ƚ', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'ᙠ', 'Ɔ', 'ᗡ', 'Ǝ', 'ᖷ', 'Ꭾ', 'H', 'I', 'Ⴑ', 'ᐴ', '⅃', 'M', 'И', 'O', 'ꟼ', 'Ọ', 'Я', 'Ƨ', 'T', 'U', 'V', 'W', 'X', 'Y', 'Ƹ'}
	for i, l := range o {
		lookup[l] = m[i]
	}
}

// Returns the boustrophedon of provided string, processed per line
// If mirrorLetters is true, all letters will be replaced with its pseudo mirrored version, which is another rune
// that resembles closely to the mirroring of the specific rune.
func ApplyToText(s string, mirrorLetters bool) string {
	apply := reverse
	if mirrorLetters {
		apply = reverseMirrored
	}

	lines := strings.Split(s, "\n")
	for i := range lines {
		if i%2 == 1 {
			lines[i] = apply(lines[i])
		}
	}
	return strings.Join(lines, "\n")
}

func reverse(s string) string {
	var reverse string
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}

func reverseMirrored(s string) string {
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
