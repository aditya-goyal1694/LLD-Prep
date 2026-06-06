// Factory code

package factory

import (
	"fmt"
	"github.com/aditya-goyal1694/LLD-Prep/design_patterns/factory/exporter"
)

func NewExporter(format string) (exporter.Exporter, error) {
    switch(format) {
        case: "pdf":
            return exporter.NewPDFExporter(), nil
        case: "csv":
            return exporter.NewCSVExporter(), nil
        case: "excel":
            return exporter.NewExcelExporter(), nil
        
        default:
            return nil, fmt.Errorf("unsupported format: %s", format)
    }
	
}