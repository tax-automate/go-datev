package datev

import (
	"archive/zip"
	"fmt"
	"github.com/google/uuid"
	"io"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

const (
	// those constants can use as 'type' or 'processID'
	bookingRelevant int = 1
	archiveRelevant int = 2

	incomingInvoices = 1
	outgoingInvoices = 2
)

// xmlData holds all information that are parsed into xml template
type xmlData struct {
	GeneratingSystem string
	Date             string
	Documents        []xmlDocument
}

type xmlDocument struct {
	GUID      uuid.UUID
	ProcessID int
	Type      int
	FileName  string
}

type XMLFactory struct {
	entries map[uuid.UUID]string
}

func NewXMLFactory() XMLFactory {
	return XMLFactory{entries: make(map[uuid.UUID]string, 0)}
}

// AddDocument add an uuid.UUID with a filePath to the entry. All entries will be written into document.xml
func (f XMLFactory) AddDocument(uid uuid.UUID, filePath string) {
	if _, ok := f.entries[uid]; !ok {
		f.entries[uid] = filePath
	}
}

// Execute create a .zip archive and copy all documents, that are related to a booking into this archive
// Also it creates the document.xml file to bind the UUIDs with the files
// Implements the DATEV XML-online interface (https://developer.datev.de/datev/platform/de/dxso)
func (f XMLFactory) Execute(saveDir string) error {
	if len(f.entries) == 0 {
		return nil
	}

	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	zipFile, err := os.OpenFile(filepath.Join(saveDir, "Belege.zip"), flags, 0644)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipw := zip.NewWriter(zipFile)
	defer zipw.Close()

	var docs []xmlDocument
	for uid, filePath := range f.entries {
		docs = append(docs, xmlDocument{
			GUID:      uid,
			ProcessID: bookingRelevant,
			Type:      outgoingInvoices,
			FileName:  filepath.Base(filePath),
		})

		err = addFileToZipArchive(filePath, zipw)
		if err != nil {
			return err
		}
	}

	tmpl, err := template.New("XML").Parse(xmlTemplate)
	if err != nil {
		return err
	}

	templateData := xmlData{
		GeneratingSystem: "tax-automate",
		Date:             time.Now().Format("2006-01-02T15:04:05"),
		Documents:        docs,
	}
	documentXML, err := zipw.Create("document.xml")
	if err != nil {
		return err
	}

	err = tmpl.Execute(documentXML, templateData)
	if err != nil {
		return err
	}

	return nil
}

func addFileToZipArchive(filePath string, w *zip.Writer) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open %s: %s", filePath, err)
	}
	defer file.Close()

	wr, err := w.Create(filepath.Base(filePath))
	if err != nil {
		return fmt.Errorf("failed to create entry for %s in zip file: %s", filepath.Base(filePath), err)
	}

	if _, err = io.Copy(wr, file); err != nil {
		return fmt.Errorf("failed to write %s to zip: %s", filePath, err.Error())
	}
	return nil
}

const xmlTemplate = `<?xml version="1.0" encoding="utf-8"?>
<archive xmlns="http://xml.datev.de/bedi/tps/document/v06.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://xml.datev.de/bedi/tps/document/v06.0 Document_v060.xsd" version="6.0"
         generatingSystem="{{ .GeneratingSystem }}">
    <header>
        <date>{{ .Date }}</date>
    </header>
    <content>{{ range .Documents }}
        <document guid="{{ .GUID }}" processID="{{ .ProcessID }}" type="{{ .Type }}">
            <extension xsi:type="File" name="{{ .FileName }}"/>
        </document>{{ end }}
    </content>
</archive>`
