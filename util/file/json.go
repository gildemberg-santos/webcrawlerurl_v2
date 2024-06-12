package file

import (
	"encoding/json"
	"os"
)

type FileJson struct {
	FileName string
	Data     interface{}
}

func NewFileJson(fileName string, data interface{}) FileJson {
	return FileJson{
		FileName: fileName,
		Data:     data,
	}
}

func (f *FileJson) Save() error {
	data_json, _ := json.MarshalIndent(f.Data, "", "  ")
	os.WriteFile(f.FileName, data_json, 0644)
	return nil
}
