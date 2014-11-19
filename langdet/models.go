package langdet

type Token struct {
	Occurrence int
	Key        string
}

type ByOccurrence []Token

func (a ByOccurrence) Len() int           { return len(a) }
func (a ByOccurrence) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByOccurrence) Less(i, j int) bool { return a[i].Occurrence < a[j].Occurrence }

type Language struct {
	Profile map[string]int
	Name    string
}

type DetectionResult struct {
	Name       string
	Distance   int
	Confidence int
}

type ResByDist []DetectionResult

func (a ResByDist) Len() int           { return len(a) }
func (a ResByDist) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ResByDist) Less(i, j int) bool { return a[i].Distance < a[j].Distance }
