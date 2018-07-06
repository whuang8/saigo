package corpus

import (
	"fmt"
	"bufio"
	"strings"
	"regexp"
	"sort"
)

var words = make([]word, 0)

func Analyze(text string) []word{
	reg, _ := regexp.Compile("[^A-Za-z']+")
	text = reg.ReplaceAllString(text, " ")
	scanner := bufio.NewScanner(strings.NewReader(strings.ToLower(text)))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		processWord(scanner.Text())
	}

	sort.Sort(ByWordCount(words))
	return words
}

func processWord(text string) {
	if len(words) == 0 {
		words = append(words, word{text: text, count: 1})
		return
	}

	for i, w := range words {
		if strings.Compare(text, w.text) == 0 {
			words[i].count++
			return
		}
	}
	words = append(words, word{text: text, count: 1})
}

func PrintWords() {
	for _, w := range words {
		fmt.Printf("%d  %s\n", w.count, w.text)
	}
}

