package main

import (
	"encoding/json"
	"fmt"
	"github.com/chrisport/go-lang-detector/langdet"
	"io/ioutil"
	"os"
)

/*
analyzed.json includes:
Arabic, English, French, German, Hebrew, Russian, Turkish
*/
func main() {

	detector := langdet.Detector{}
	// Analyze different languages from files and and write to analyzed.json:
	detector.AddLanguage(GetTextFromFile("samples/english.txt"), "english")
	detector.AddLanguage(GetTextFromFile("samples/german.txt"), "german")
	detector.AddLanguage(GetTextFromFile("samples/french.txt"), "french")
	detector.AddLanguage(GetTextFromFile("samples/turkish.txt"), "turkish")
	detector.AddLanguage(GetTextFromFile("samples/arabic"), "arabic")
	detector.AddLanguage(GetTextFromFile("samples/hebrew"), "hebrew")
	detector.AddLanguage(GetTextFromFile("samples/russian"), "russian")
	bytes, _ := json.Marshal(*detector.Languages)
	WriteToFile(bytes, "analyzed.json")

	//detector := langdet.NewDefaultDetector()
	testString := GetTextFromFile("example_input.txt")
	result := detector.GetClosestLanguage(testString)
	fmt.Println("GetClosestLanguage returns:\n", "    ", result)

	fullResults := detector.GetLanguages(testString)
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
