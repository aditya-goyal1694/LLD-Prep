// Factory code

package factory

import (
	"fmt"

	"lld/design_patterns/factory/exporter"
)

func NewExporter(format string) (exporter.Exporter, error) {
	switch format {
	case "pdf":
		return exporter.NewPDFExporter(), nil
	case "csv":
		return exporter.NewCSVExporter(), nil
	case "excel":
		return exporter.NewExcelExporter(), nil
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}
}
