package langdet

import (
	"unicode"
	"fmt"
	"unicode/utf8"
)

type CJKLanguageComparator struct {
}

func (cl *CJKLanguageComparator) GetName() string {
	return "Chinese"
}

func (cl *CJKLanguageComparator) CompareTo(_ func() map[string]int, originalInput string) DetectionResult {
	fl := utf8.RuneCountInString(originalInput)
	cc := 0
	for _, r := range originalInput {
		if unicode.Is(unicode.Han, r) {
			cc++
		}
	}
	r := (float64(cc) / float64(fl)) * 100
	return DetectionResult{"Chinese", int(r)}
}
