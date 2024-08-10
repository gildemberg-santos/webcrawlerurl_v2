package main

import (
	"log"
	"os"

	downloaddrive "github.com/gildemberg-santos/webcrawlerurl_v2/util/download_drive"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	chromeDownloadDriver := os.Getenv("CHROME_DOWNLOAD_DRIVER")
	chromePath := os.Getenv("CHROME_PATH")

	if _, err := os.Stat(chromePath); os.IsNotExist(err) {
		log.Default().Printf("O chromedriver não foi encontrado no caminho especificado: %s", chromePath)
		err = downloaddrive.NewDownloadDrive(chromeDownloadDriver, chromePath).Call()
		if err != nil {
			panic(err)
		}

		err = os.Chmod(chromePath, 0755)
		if err != nil {
			panic(err)
		}
	}
}
