package main

import (
	"fmt"
	"github.com/chrisport/go-lang-detector/langdet"

	"io/ioutil"
	"os"
)

func main() {
	detector := langdet.NewDefaultLanguages()

	result := detector.GetClosestLanguage("ont ne comprend rien")
	fmt.Println("GetClosestLanguage returns:\n", "    ", result)

	fullResults := detector.GetLanguages("ont ne comprend rien")
	fmt.Println("GetLanguages returns:")
	for _, r := range fullResults {
		fmt.Println("    ", r.Name, r.Confidence, "%")
	}

}

// GetTextFromFile returns the content of file (identified by given fileName) as text
func GetTextFromFile(fileName string) string {
	text, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(text)
}

// WriteToFile writes a content into a file with specified name
func WriteToFile(content []byte, fileName string) {
	err := ioutil.WriteFile(fileName, content, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
