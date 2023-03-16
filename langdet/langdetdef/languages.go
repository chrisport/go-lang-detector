package langdetdef

import (
	"encoding/json"
	"fmt"
	"unicode"

	"github.com/chrisport/go-lang-detector/langdet"
	"github.com/chrisport/go-lang-detector/langdet/internal"
)

func init() {
	lan := []langdet.Language{}

	if err := json.Unmarshal(internal.DefaultLanguageDefs, &lan); err != nil {
		panic(fmt.Sprintf("unable to initialize default languages - corrupt embedded asset: %v", err))
	}

	for i := range lan {
		switch lan[i].Name {
		case "russian":
			RUSSIAN = &lan[i]
		case "french":
			FRENCH = &lan[i]
		case "english":
			ENGLISH = &lan[i]
		case "turkish":
			TURKISH = &lan[i]
		case "german":
			GERMAN = &lan[i]
		case "hebrew":
			HEBREW = &lan[i]
		case "arabic":
			ARABIC = &lan[i]
		}
		defaultLanguages[lan[i].Name] = &lan[i]
	}
}

var CHINESE_JAPANESE_KOREAN = &langdet.UnicodeRangeLanguageComparator{"CJK", unicode.Han}
var HEBREW langdet.LanguageComparator
var ARABIC langdet.LanguageComparator
var ENGLISH langdet.LanguageComparator
var RUSSIAN langdet.LanguageComparator
var GERMAN langdet.LanguageComparator
var FRENCH langdet.LanguageComparator
var TURKISH langdet.LanguageComparator

func DefaultLanguages() []langdet.LanguageComparator {
	return []langdet.LanguageComparator{CHINESE_JAPANESE_KOREAN, HEBREW, ARABIC, ENGLISH, RUSSIAN, GERMAN, FRENCH, TURKISH}
}

// NewWithDefaultLanguages returns a new Detector with the default languages, if loaded:
// currently: Arabic, English, French, German, Hebrew, Russian, Turkish, Chinese
func NewWithDefaultLanguages() langdet.Detector {
	return langdet.Detector{Languages: DefaultLanguages(),
		MinimumConfidence: langdet.DefaultMinimumConfidence,
		NDepth:            langdet.DEFAULT_NDEPTH}
}

var defaultLanguages = make(map[string]langdet.LanguageComparator)
