package unzipdrive_test

import (
	"io"
	"os"
	"testing"

	unzipdrive "github.com/gildemberg-santos/webcrawlerurl_v2/util/unzip_drive"
	"github.com/stretchr/testify/assert"
)

func TestUnzipDrive_Call(t *testing.T) {
	arquivoOrigem, _ := os.Open("TestTMP.zip")
	defer arquivoOrigem.Close()
	arquivoDestino, _ := os.Create("Test.zip")
	defer arquivoDestino.Close()
	io.Copy(arquivoDestino, arquivoOrigem)

	unzipDrive := unzipdrive.NewUnzipDrive("Test.zip", "Test")

	assert.Nil(t, unzipDrive.Call())

	os.RemoveAll("Test")
}
