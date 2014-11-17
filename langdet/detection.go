package langdet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const nDepth = 3

var defaultLanguages []Language

func init() {
	analyzedInput, err := ioutil.ReadFile("analyzed.json")
	if err != nil {
		fmt.Println("go-lang-detector/langdet: Could not read default languages from analyzed.json")
		return
	}
	err = json.Unmarshal(analyzedInput, &defaultLanguages)
	if err != nil {
		fmt.Println("go-lang-detector/langdet: Could not unmarshall default languages from analyzed.json")
		return
	}
}

type Detector struct {
	Languages *[]Language
}

func NewDefaultDetector() Detector {
	defaultCopy := make([]Language, len(defaultLanguages))
	copy(defaultCopy, defaultLanguages)
	return Detector{&defaultCopy}
}

func (this *Detector) AddLanguage(textToAnalyze, languageName string) {
	if this.Languages == nil {
		newSlice := make([]Language, 0, 0)
		this.Languages = &newSlice
	}
	analyzedLanguage := Analyze(textToAnalyze, languageName)
	updatedList := append(*this.Languages, analyzedLanguage)
	*this.Languages = updatedList
}

func (this *Detector) GetClosestLanguage(text string) string {
	occurrenceMap := createOccurenceMap(text, nDepth)
	lookupMap := createRankLookupMap(occurrenceMap)
	result := this.closestFromTable(lookupMap)
	return result
}

func (this *Detector) closestFromTable(lookupMap map[string]int) string {
	minDist := 999999999
	minLan := "undefined"
	for _, language := range *this.Languages {
		dist := getDistance(lookupMap, language.Profile)
		fmt.Println(language.Name, ":", dist)
		if dist < minDist {
			minDist = dist
			minLan = language.Name
		}
	}
	return minLan
}

func getDistance(mapA, mapB map[string]int) int {
	var result int
	for key, rankA := range mapA {
		if rankA > 300 {
			continue
		}
		if rankB, ok := mapB[key]; ok {
			diff := rankB - rankA
			if diff < 0 {
				diff = diff * (-1)
			}
			result += diff
		} else {
			result += 300
		}
	}
	return result
}
