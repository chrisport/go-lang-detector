package langdet

import (
	"bytes"
	"sort"
	"strings"
	"unicode/utf8"
)

func analyze(text, name string) Language {
	theMap := createOccurenceMap(text, nDepth)
	ranked := createRankLookupMap(theMap)
	return Language{Name: name, Profile: ranked}
}

func createRankLookupMap(input map[string]int) map[string]int {
	tokens := make([]Token, len(input))
	counter := 0
	for k, v := range input {
		tokens[counter] = Token{Key: k, Occurrence: v}
		counter++
	}
	sort.Sort(ByOccurrence(tokens))
	result := make(map[string]int)
	length := len(tokens)
	for i, curr := range tokens {
		result[curr.Key] = length - i
	}
	return result
}

func createOccurenceMap(text string, gramDepth int) map[string]int {
	text = cleanText(text)
	tokens := strings.Split(text, " ")
	result := make(map[string]int)
	for _, token := range tokens {
		analyseToken(result, token, gramDepth)
	}
	return result
}

func analyseToken(resultMap map[string]int, token string, gramDepth int) {
	if len(token) == 0 {
		return
	}
	for i := 1; i <= gramDepth+1; i++ {
		generateNthGrams(resultMap, token, i)
	}
}

func generateNthGrams(resultMap map[string]int, text string, n int) {
	padding := createPadding(n - 1)
	text = padding + text + padding
	upperBound := utf8.RuneCountInString(text) - (n - 1)
	for p := 0; p < upperBound; p++ {
		currentToken := text[p : p+n]
		resultMap[currentToken] += 1
	}
}

func createPadding(length int) string {
	var buffer bytes.Buffer
	padding := "_"
	for i := 0; i < length; i++ {
		buffer.WriteString(padding)
	}
	return buffer.String()
}

func cleanText(text string) string {
	text = strings.Replace(text, "\n", " ", -1)
	text = strings.Replace(text, ",", " ", -1)
	text = strings.Replace(text, "#", " ", -1)
	text = strings.Replace(text, "/", " ", -1)
	text = strings.Replace(text, "\\", " ", -1)
	text = strings.Replace(text, ".", " ", -1)
	text = strings.Replace(text, "!", " ", -1)
	text = strings.Replace(text, "?", " ", -1)
	text = strings.Replace(text, ":", " ", -1)
	text = strings.Replace(text, ";", " ", -1)
	text = strings.Replace(text, "-", " ", -1)
	text = strings.Replace(text, "'", " ", -1)
	text = strings.Replace(text, "\"", " ", -1)
	text = strings.Replace(text, "_", " ", -1)
	text = strings.Replace(text, "*", " ", -1)
	text = strings.Replace(text, "1", "", -1)
	text = strings.Replace(text, "2", "", -1)
	text = strings.Replace(text, "3", "", -1)
	text = strings.Replace(text, "4", "", -1)
	text = strings.Replace(text, "5", "", -1)
	text = strings.Replace(text, "6", "", -1)
	text = strings.Replace(text, "7", "", -1)
	text = strings.Replace(text, "8", "", -1)
	text = strings.Replace(text, "9", "", -1)
	text = strings.Replace(text, "0", "", -1)
	text = strings.Replace(text, "  ", " ", -1)
	return text
}
