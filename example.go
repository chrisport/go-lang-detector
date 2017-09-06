package main

import (
	"fmt"
	"github.com/chrisport/go-lang-detector/langdet/langdetdef"
)

func main() {
	detector := langdetdef.NewWithDefaultLanguages()

	result := detector.GetClosestLanguage("ont ne comprend rien")
	fmt.Println("GetClosestLanguage returns:\n", "    ", result)

	fullResults := detector.GetLanguages("ont ne comprend rien")
	fmt.Println("GetLanguages returns:")
	for _, r := range fullResults {
		fmt.Println("    ", r.Name, r.Confidence, "%")
	}

	fullResults = detector.GetLanguages("义勇军进行曲")
	fmt.Println("GetLanguages for Chinese returns:")
	for _, r := range fullResults {
		fmt.Println("    ", r.Name, r.Confidence, "%")
	}

}
