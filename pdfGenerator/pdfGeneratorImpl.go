package pdfGenerator

import (
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/google/uuid"
	"os"
)

type wk struct {
	rootPath string
}

func NewWKHtmlToPDF(rootPath string) PDFGeneratorInterface {
	return &wk{rootPath: rootPath}
}

func (wk *wk) Create(htmlFile string) (string, error) {
	file, err := os.Open(htmlFile)
	if err != nil {
		return "", err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	pdfg.AddPage(wkhtmltopdf.NewPageReader(file))

	if err := pdfg.Create(); err != nil {
		return "", err
	}

	pdfName := wk.rootPath + "/" + uuid.New().String() + ".pdf"

	if err := pdfg.WriteFile(pdfName); err != nil {
		return "", err
	}
	return pdfName, nil
}
