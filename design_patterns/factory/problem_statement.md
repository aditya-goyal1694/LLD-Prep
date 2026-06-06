## You are building a Document Export System.

The application allows users to export reports in different formats.

Current supported formats:

PDF
CSV
Excel

Every exporter should support:

```Go
Export(data string) error
```

Expected behavior:

PDF Exporter    -> Generates PDF
CSV Exporter    -> Generates CSV
Excel Exporter  -> Generates Excel


## New Requirement

Product team says:

Next month we may support:

JSON
XML
Google Sheets

and more formats later.

Developers should be able to add new exporters without modifying existing business logic.