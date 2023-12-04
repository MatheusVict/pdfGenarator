package main

import (
	"fmt"
	"log"
	"pdfGenarator/htmlParser"
	"pdfGenarator/pdfGenerator"
)

type Data struct {
	Name string
}

func main() {
	html := htmlParser.New("tmp")
	wk := pdfGenerator.NewWKHtmlToPDF("tmp")

	dataHTML := Data{Name: "Matheus"}
	htmlGenerator, err := html.Create("templates/example.html", dataHTML)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("html ", htmlGenerator)

	filePdfName, err := wk.Create(htmlGenerator)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("pdf gerade", filePdfName)
}
