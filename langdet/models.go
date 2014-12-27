package langdet

// Token represents a text token and its occurence in an analyzed text
type Token struct {
	Occurrence int
	Key        string
}

// ByOccurrence represents an array of tokens which can be sorted by occurrences of the tokens.
type ByOccurrence []Token

func (a ByOccurrence) Len() int      { return len(a) }
func (a ByOccurrence) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByOccurrence) Less(i, j int) bool {
	if a[i].Occurrence == a[j].Occurrence {
		return a[i].Key < a[i].Key
	}
	return a[i].Occurrence < a[j].Occurrence
}

// Language represents a language by its name and the profile ( map[token]OccurrenceRank )
type Language struct {
	Profile map[string]int
	Name    string
}

// DetectionResult represents the result from comparing 2 Profiles. It includes the confidence which is basically the
// the relative distance between the two profiles.
type DetectionResult struct {
	Name       string
	Confidence int
}

//ResByConf represents an array of DetectionResult and can be sorted by Confidence.
type ResByConf []DetectionResult

func (a ResByConf) Len() int           { return len(a) }
func (a ResByConf) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ResByConf) Less(i, j int) bool { return a[i].Confidence > a[j].Confidence }
