package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_scoreNumberWords(t *testing.T) {
	score := scoreNumberWords("test01 test02 test03 test04 test05", 5, 10)
	assert.Equal(t, float32(10), score)
}
