package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChackAI_Call(t *testing.T) {
	chackTeste := []string{"teste01", "teste02", "teste03", "teste04", "teste05"}
	chackgpt3 := NewChackAi(5, "teste01 teste02 teste03 teste04 teste05 teste06")
	chackgpt3.Call()

	assert.Equal(t, chackTeste, chackgpt3.ListChacks)
}
