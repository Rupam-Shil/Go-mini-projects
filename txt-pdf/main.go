package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

type TextToPDF struct{
	TextPath string
	PdfPath string
}

func main() {
	fmt.Println("Welcome to text to pdf")
	fmt.Println("Please enter the file name")

	reader := bufio.NewReader(os.Stdin)
	fileName, err := reader.ReadString('\n')
	handleError(err)
	fileName = strings.TrimSpace(fileName)
	details :=  TextToPDF{fileName+".txt", fileName+".pdf"}
	err = details.CreatePdf()
	handleError(err)
	content, err := ioutil.ReadFile(details.TextPath)
	handleError(err)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial","B",14)
	pdf.MultiCell(190, 5, string(content), "0","0", false)

	_ = pdf.OutputFileAndClose(details.PdfPath)
	fmt.Println("PDF created")
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (t TextToPDF) CreatePdf() error {
	fmt.Println(t.TextPath)
	content, err := ioutil.ReadFile(t.TextPath)
	handleError(err)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial","B",14)
	pdf.MultiCell(190, 5, string(content), "0","0", false)

	return pdf.OutputFileAndClose(t.PdfPath)
}