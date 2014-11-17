package main

import (
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

	/*
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
	*/
	detector := langdet.NewDefaultDetector()
	testString := GetTextFromFile("example.txt")
	result := detector.GetClosestLanguage(testString)
	fmt.Println(result)
}

func GetTextFromFile(fileName string) string {
	text, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(text)
}

func WriteToFile(content []byte, name string) {
	err := ioutil.WriteFile(name+"out.txt", content, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
