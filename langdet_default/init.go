package langdet_default

import (
	"github.com/chrisport/go-lang-detector/langdet"
	"bytes"
	"encoding/json"
	"fmt"
)

var DefaultLanguages = []langdet.Language{}

func init() {
	langdet.InitWithDefaultFromReader(bytes.NewReader(_default_languagesJson))
	err := json.Unmarshal(_default_languagesJson, &DefaultLanguages)
	if err != nil {
		fmt.Println("Could not initialize default languages")
	}
}
