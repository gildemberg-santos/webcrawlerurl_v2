package downloaddrive

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	unzipdrive "github.com/gildemberg-santos/webcrawlerurl_v2/util/unzip_drive"
)

type DownloadDrive struct {
	UrlLocation string
	PathDrives  string
}

func NewDownloadDrive(url string, pathDrives string) *DownloadDrive {
	log.Printf("Initializing download drive from %s to %s", url, pathDrives)
	return &DownloadDrive{
		UrlLocation: url,
		PathDrives:  pathDrives,
	}
}

func (d *DownloadDrive) Call() error {
	filename := path.Base(d.UrlLocation)
	drives := strings.Split(d.PathDrives, "/")[1]

	log.Printf("Downloading %s to %s", d.UrlLocation, filename)
	resp, err := http.Get(d.UrlLocation)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	src := "./" + filename
	dest := "./" + drives

	err = unzipdrive.NewUnzipDrive(src, dest).Call()
	return err
}
