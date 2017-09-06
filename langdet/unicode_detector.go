package langdet

import (
	"unicode/utf8"
	"unicode"
)

// Chinese, Japanese, Korean Charset
type UnicodeRangeLanguageComparator struct {
	Name       string
	RangeTable *unicode.RangeTable
}

func (u *UnicodeRangeLanguageComparator) GetName() string {
	return u.Name
}

func (u *UnicodeRangeLanguageComparator) CompareTo(_ func() map[string]int, originalInput string) DetectionResult {
	fl := utf8.RuneCountInString(originalInput)
	cc := 0
	for _, r := range originalInput {
		if unicode.Is(u.RangeTable, r) {
			cc++
		}
	}
	r := (float64(cc) / float64(fl)) * 100
	return DetectionResult{u.Name, int(r)}
}
