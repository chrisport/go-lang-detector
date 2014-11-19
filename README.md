# Language Detector
This golang library provides functionality to analyze and recognize text.
It is an implementation of:
N-Gram-Based Text Categorization
William B. Cavnar and John M. Trenkle
Environmental Research Institute of Michigan P.O. Box 134001
Ann Arbor MI 48113-4001

### Language profile
A language profile is a ```map[string] int```that maps n-gram tokens to its occurrency-rank. So for the most
frequent token 'X' of the analyzed text, map['X'] will be 1.

## Usage
### Detect
The default detector supports the following languages:
**Arabic, English, French, German, Hebrew, Russian, Turkish**

```
    detector := langdet.NewDefaultDetector()
	testString := GetTextFromFile("example_input.txt")
	result := detector.GetClosestLanguage(testString)
	fmt.Println(result)

output:
    'english'
```

### Analyze
The result will be a Language object, containing the specified name and the profile
example:

``` 
    language := langdet.Analyze(text_sample,"french")
    language.Profile // the profile
    language.Name // the name
```

### Add more languages
New languages can directly be analyzed and added to a detector by providing a text sample:

```
    text_sample := GetTextFromFile("samples/polish.txt")
    detector.AddLanguage(text_sample, "polish")
```

The text sample should be bigger then 200kb and can be "dirty" (special chars, lists, etc.), but the language
should not change for long parts.