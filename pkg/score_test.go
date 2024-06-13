package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_scoreNumberWords is a test function for the scoreNumberWords function.
//
// It tests the scoreNumberWords function by passing a string of words, a limit
// for the number of words, and a note value. It asserts that the calculated
// score is equal to the expected value.
//
// Parameters:
// - t: The testing.T object for running the test.
//
// Return type: None.
func Test_scoreNumberWords(t *testing.T) {
	score := scoreNumberWords("test01 test02 test03 test04 test05", 5, 10)
	assert.Equal(t, float32(10), score)
}
