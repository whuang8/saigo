package corpus

type word struct {
	text  string
	count int
}

type ByWordCount []word


func (w ByWordCount) Len() int {
	return len(w)
}

func (w ByWordCount) Less(i, j int) bool {
	return w[i].count > w[j].count
}

func (w ByWordCount) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}
