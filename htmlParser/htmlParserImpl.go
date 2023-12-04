package htmlParser

import (
	"github.com/google/uuid"
	"os"
	"text/template"
)

type htmlStruct struct {
	rootPath string
}

func New(rootPath string) HTMLParserInterface {
	return &htmlStruct{rootPath: rootPath}
}

func (h *htmlStruct) Create(templateName string, data interface{}) (string, error) {
	templateWriter, err := template.ParseFiles(templateName)
	if err != nil {
		return "", err
	}

	fileName := h.rootPath + "/" + uuid.New().String() + ".html"

	fileWriter, err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	if err := templateWriter.Execute(fileWriter, data); err != nil {
		return fileName, err
	}
	return fileName, nil
}
