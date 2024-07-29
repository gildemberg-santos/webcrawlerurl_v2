package unzipdrive

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type UnzipDrive struct {
	Source    string
	PathDrive string
}

func NewUnzipDrive(source, pathDrive string) *UnzipDrive {
	log.Printf("Initializing UnzipDrive with source %s and path %s", source, pathDrive)
	return &UnzipDrive{
		Source:    source,
		PathDrive: pathDrive,
	}
}

func (u *UnzipDrive) Call() error {
	log.Printf("Unzipping %s to %s", u.Source, u.PathDrive)
	r, err := zip.OpenReader(u.Source)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(u.PathDrive, f.Name)

		if !strings.HasPrefix(fpath, filepath.Clean(u.PathDrive)+string(os.PathSeparator)) {
			return fmt.Errorf("arquivo fora do diretório alvo: %s", fpath)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer outFile.Close()

		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		_, err = io.Copy(outFile, rc)
		if err != nil {
			return err
		}
	}

	os.Remove(u.Source)
	log.Printf("Unzipped %s to %s", u.Source, u.PathDrive)
	return nil
}
