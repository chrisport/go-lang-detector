package langdet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
)

const nDepth = 3

var MaxDistanceInPercentage float64 = 0.7

var defaultLanguages []Language

func init() {
	analyzedInput, err := ioutil.ReadFile("default.json")
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
	results := this.closestFromTable(lookupMap)

	maxPossibleDistance := len(lookupMap) * 300
	maxAllowedDistance := int(MaxDistanceInPercentage * float64(maxPossibleDistance))

	if results[0].Distance < maxAllowedDistance {
		return results[0].Name
	}
	return "undefined"
}

func (this *Detector) GetLanguages(text string) []DetectionResult {
	occurrenceMap := createOccurenceMap(text, nDepth)
	lookupMap := createRankLookupMap(occurrenceMap)
	results := this.closestFromTable(lookupMap)
	return results
}

func (this *Detector) closestFromTable(lookupMap map[string]int) []DetectionResult {
	results := []DetectionResult{}
	maxPossibleDistance := len(lookupMap) * 300
	for _, language := range *this.Languages {
		dist := getDistance(lookupMap, language.Profile)
		relativeDistance := 1 - float64(dist)/float64(maxPossibleDistance)
		confidence := int(relativeDistance * 100)
		results = append(results, DetectionResult{Name: language.Name, Distance: dist, Confidence: confidence})
	}

	sort.Sort(ResByDist(results))
	return results
}

func getDistance(mapA, mapB map[string]int) int {
	var result int
	for key, rankA := range mapA {
		if rankA > 300 {
			continue
		}
		var diff int
		if rankB, ok := mapB[key]; ok {
			diff = rankB - rankA

			if diff > 300 || diff < -300 {
				diff = 300
			} else if diff < 0 {
				diff = diff * (-1)
			}
		} else {
			diff = 300
		}
		result += diff

	}
	return result
}
