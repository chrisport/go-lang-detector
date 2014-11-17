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
