package corpus

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"sort"
)

func TestCorpus(t *testing.T) {
	result := Analyze("Are you serious? I knew you were.")
	assert.Equal(t, len(result), 6)
	assert.Equal(t, result[0].text, "you")
	assert.Equal(t, result[0].count, 2)
}

func TestWordSort(t *testing.T) {
	w1 := word{text: "word1", count: 1}
	w2 := word{text: "word2", count: 2}
	w3 := word{text: "word3", count: 3}

	words = append(words, w1, w2, w3)
	sort.Sort(ByWordCount(words))
	assert.Equal(t, words[0].text, "word3")
	assert.Equal(t, words[0].count, 3)
}
