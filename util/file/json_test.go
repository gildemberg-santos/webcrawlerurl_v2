package file_test

import (
	"os"
	"testing"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/file"
	"github.com/stretchr/testify/assert"
)

func TestFile_Call(t *testing.T) {
	file := file.NewFileJson("test.json", map[string]string{"teste": "teste"})
	err := file.Save()
	assert.Nil(t, err)
	assert.FileExists(t, "test.json")
	os.Remove("test.json")
}
