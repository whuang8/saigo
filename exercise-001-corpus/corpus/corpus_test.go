package corpus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCorpus(t *testing.T) {
	result := Analyze("William william, is Is, MY my mY, NAME. name. NaMe!!")
	assert.Equal(t, len(result), 4)
	assert.Equal(t, result["william"], 2)
	assert.Equal(t, result["is"], 2)
	assert.Equal(t, result["my"], 3)
	assert.Equal(t, result["name"], 3)
}
