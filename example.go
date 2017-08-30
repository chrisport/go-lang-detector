package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/chrisport/go-lang-detector/langdet"
)

/*
analyzed.json includes:
Arabic, English, French, German, Hebrew, Russian, Turkish
*/
func main() {

	//sample using Reader to Initialize default languages
	//	analyzedInput, _ := ioutil.ReadFile("default_languages2.json")
	//	s := string(analyzedInput[:1652088])
	//	langdet.InitWithDefaultFromReader(strings.NewReader(s))
	//	detector := langdet.NewDefaultLanguages()

	//sample by manually analyzing languages
	//	 Analyze different languages from files and and write to analyzed.json:
	detector := langdet.Detector{}
	detector.AddLanguageFromText(GetTextFromFile("samples/english.txt"), "english")
	detector.AddLanguageFromText(GetTextFromFile("samples/german.txt"), "german")
	detector.AddLanguageFromText(GetTextFromFile("samples/french.txt"), "french")
	detector.AddLanguageFromText(GetTextFromFile("samples/turkish.txt"), "turkish")
	detector.AddLanguageFromText(GetTextFromFile("samples/arabic"), "arabic")
	detector.AddLanguageFromText(GetTextFromFile("samples/hebrew"), "hebrew")
	detector.AddLanguageFromText(GetTextFromFile("samples/russian"), "russian")
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
