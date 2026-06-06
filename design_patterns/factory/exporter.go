// App code

package exporter

import (
    "fmt"
)

type Exporter interface {
	Export(data string) error
}


type PDFExporter struct {}

func NewPDFExporter() *PDFExporter {
    return &PDFExporter{}
}

func (p *PDFExporter) Export(data string) error {
    fmt.Printf("Exporting pdf file...")
    return nil
}


type CSVExporter struct {}

func NewCSVExporter() *CSVExporter {
    return &CSVExporter{}
}

func (c *CSVExporter) Export(data string) error {
    fmt.Printf("Exporting csv file...")
    return nil
}


type ExcelExporter struct {}

func NewExcelExporter() *ExcelExporter {
    return &ExcelExporter{}
}

func (e *ExcelExporter) Export(data string) error {
    fmt.Printf("Exporting excel file...")
    return nil
}
