package main

import (
	"fmt"
	"github.com/chrisport/go-lang-detector/langdet"
)

func main() {
	detector := langdet.NewDefaultLanguages()

	result := detector.GetClosestLanguage("Comment ça va? Est-ce que ça va bien?")
	fmt.Println("GetClosestLanguage returns:\n", "    ", result)

	fullResults := detector.GetLanguages("Comment ça va? Est-ce que ça va bien?")
	fmt.Println("GetLanguages returns:")
	for _, r := range fullResults {
		fmt.Println("    ", r.Name, r.Confidence, "%")
	}

	detector.AddLanguageComparators(&langdet.CJKLanguageComparator{})

	detector.GetLanguages("ont ne comprend rien")
	detector.GetLanguages("义勇军进行曲")
	fmt.Println("GetLanguages returns:")
	for _, r := range fullResults {
		fmt.Println("    ", r.Name, r.Confidence, "%")
	}
}
