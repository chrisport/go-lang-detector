package langdet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
)

const nDepth = 4

// MinimumConfidence is the minimum confidence that a language-match must have to be returned as detected language
var MinimumConfidence float32 = 0.7

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

	if results[0].Confidence < asPercent(MinimumConfidence) {
		return "undefined"
	}
	return results[0].Name
}

func (this *Detector) GetLanguages(text string) []DetectionResult {
	occurrenceMap := createOccurenceMap(text, nDepth)
	lookupMap := createRankLookupMap(occurrenceMap)
	results := this.closestFromTable(lookupMap)
	return results
}

func (this *Detector) closestFromTable(lookupMap map[string]int) []DetectionResult {
	results := []DetectionResult{}
	inputSize := len(lookupMap)
	if inputSize > 300 {
		inputSize = 300
	}
	for _, language := range *this.Languages {
		lSize := len(language.Profile)
		maxPossibleDistance := lSize * inputSize
		dist := getDistance(lookupMap, language.Profile, lSize)
		relativeDistance := 1 - float64(dist)/float64(maxPossibleDistance)
		confidence := int(relativeDistance * 100)
		results = append(results, DetectionResult{Name: language.Name, Confidence: confidence})
	}

	sort.Sort(ResByConf(results))
	return results
}

// getDistance calculates the out-of-place distance between two Profiles,
// taking into account only items of mapA, that have a value bigger then 300
func getDistance(mapA, mapB map[string]int, maxDist int) int {
	var result int
	negMaxDist := ((-1) * maxDist)
	for key, rankA := range mapA {
		if rankA > 300 {
			continue
		}
		var diff int
		if rankB, ok := mapB[key]; ok {
			diff = rankB - rankA

			if diff > maxDist || diff < negMaxDist {
				diff = maxDist
			} else if diff < 0 {
				diff = diff * (-1)
			}
		} else {
			diff = maxDist
		}
		result += diff
	}
	return result
}

func asPercent(input float32) int {
	return int(input * 100)
}
