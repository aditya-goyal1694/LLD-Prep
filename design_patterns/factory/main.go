// Client code

package main

import (
	"log"

	"lld/design_patterns/factory/factory"
)

func main() {
	exporter, err := factory.NewExporter("pdf")
	if err != nil {
		log.Fatal(err)
	}

	data := "This will be exported."
	exporter.Export(data)
}