[![wercker status](https://app.wercker.com/status/9e2a695f35c1cf5e1cac46035e4ca7a6/m/ "wercker status")](https://app.wercker.com/project/byKey/9e2a695f35c1cf5e1cac46035e4ca7a6)
[![Coverage Status](https://img.shields.io/coveralls/chrisport/go-lang-detector.svg)](https://coveralls.io/r/chrisport/go-lang-detector?branch=master)

Breaking changes in v0.2: please see chapter "Migration" below.
Previous version is available under Release v0.1: https://github.com/chrisport/go-lang-detector/releases/tag/v0.1

# Language Detector

This golang library provides functionality to analyze and recognize language based on text.

The implementation is based on the following paper:  
N-Gram-Based Text Categorization  
William B. Cavnar and John M. Trenkle  
Environmental Research Institute of Michigan P.O. Box 134001  
Ann Arbor MI 48113-4001

### Detection by Language profile
A language profile is a ```map[string] int```that maps n-gram tokens to its occurrency-rank. So for the most
frequent token 'X' of the analyzed text, map['X'] will be 1.

### Detection by unicode range
A second way to detect the language is by the unicode range used in the text.
Golang has a set of predefined unicode ranges in package unicode, which can be used
easily, for example for detecting Chinese/Japanese/Korean:
``` go
var CHINESE_JAPANESE_KOREAN = &langdet.UnicodeRangeLanguageComparator{"CJK", unicode.Han}
```
## Usage
### Detect
#### Get the closest language:
The default detector supports the following languages:
**Arabic, English, French, German, Hebrew, Russian, Turkish**

``` go
    detector := langdetdef.NewWithDefaultLanguages()
	testString := "do not care about quantity"
	result := detector.GetClosestLanguage(testString)
	fmt.Println(result)

output:
    english
```
by setting the value langdet.MinimumConfidence (0-1), you can set the accepted confidence level.
E.g. 0.7 --> if langdet is 70% or higher sure that the language matches, return it, else it returns 'undefined'

#### Get Language Probabilities
GetClosestLanguage will return the language that most probably matches. To get the result of all analyzed language, you can use
GetLanguage, which will return you all analyzed languages and their percentage of matching the input snippet

 ```
 testString := "ont permis d'identifier"
 GetLanguages returns:
     french 86 %
     english 79 %
     german 71 %
     turkish 54 %
     hebrew 39 %
     arabic 8 %
     russian 5 %


 ```

### Analyze new language

For analysing a new language random Wikipedia articles in the target languages are ideal. The result will be a Language object, containing the specified name and the profile
example:

``` go
    language := langdet.Analyze(text_sample, "french")
    language.Profile // language profile in form of map[string]int as defined above
    language.Name // the name that was given as parameter
```

### Add more languages
New languages can directly be analyzed and added to a detector by providing a text sample:

``` go
    text_sample := GetTextFromFile("samples/polish.txt")
    detector.AddLanguageFrom(text_sample, "polish")
```

The text sample should be bigger then 200kb and can be "dirty" (special chars, lists, etc.), but the language
should not change for long parts.

Alternatively Analyze can be used and the resulting language can added using AddLanguage method:

``` go
    text_sample := GetTextFromFile("samples/polish.txt")
    french := langdet.Analyze(text_sample, "french")

    //language can be added selectively to detectors
    detectorA.AddLanguage(french)
    detectorC.AddLanguage(french)
```

## Migration to v0.2

This library has been adapted to a more convenient and more idiomatic way.
- Default languages are provided in Go code and there is no need for adding the json file anymore.
- All code related to defaults has been moved to package langdetdef
- Default languages can be added using the provided interfaces:
``` go
// detector with default languages
detector := langdetdef.NewWithDefaultLanguages()

// add all to existing detector
defaults := langdetdef.DefaultLanguages()
detector.AddLanguageComparators(defaults...)

// add selectively
detector.AddLanguageComparators(langdetdef.CHINESE_JAPANESE_KOREAN, langdetdef.ENGLISH)
```
- InitWithDefaultFromXY has been removed, custom default languages can be unmarshaled manually and added to a detector through
AddLanguage interface:
```
detector := langdet.NewDetector()
customLanguages := []langdet.Language{}

_ = json.Unmarshal(bytesFromFile, &customLanguages)
detector.AddLanguage(customLanguages...)
```

## Contribution

Suggestions and Bug reports can be made through Github issues.
Contributions are welcomed, there is currently no need to open an issue for it, but please follow the code style, including descriptive tests with [GoConvey](http://goconvey.co/).

## License

Licensed under [Apache 2.0](LICENSE).
