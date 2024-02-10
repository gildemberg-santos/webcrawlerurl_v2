package chunck_test

import (
	"testing"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/chunck"
	"github.com/stretchr/testify/assert"
)

func TestChunck_Call(t *testing.T) {
	chunckTeste := []string{"teste01", "teste02", "teste03", "teste04", "teste05"}
	chunckgpt3 := chunck.NewChunck(5, "teste01 teste02 teste03 teste04 teste05 teste06")

	assert.Equal(t, chunckTeste, chunckgpt3.Call().ListChuncks)
}
