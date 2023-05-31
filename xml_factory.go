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

// documentEntries stores the filepath of a document as value and the uuid as key
type documentEntries map[uuid.UUID]string

// add adds new entry to map if uuid not exists
func (d documentEntries) add(k uuid.UUID, v string) {
	if _, ok := d[k]; !ok {
		d[k] = v
	}
}

type xmlFactory struct {
	entries documentEntries
}

func newXMLFactory() xmlFactory {
	return xmlFactory{entries: make(documentEntries, 0)}
}

// Execute create a .zip archive and copy all documents, that are related to a booking into this archive
// Also it creates the document.xml file to bind the UUIDs with the files
// Implements the DATEV XML-online interface (https://developer.datev.de/datev/platform/de/dxso)
func (f xmlFactory) Execute(saveDir string) error {
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
			Type:      bookingRelevant,
			FileName:  filepath.Base(filePath),
		})

		err = addFileToZipArchive(filePath, zipw)
		if err != nil {
			return err
		}
	}

	tmpl, err := template.ParseFiles("./xml_datev.xml")
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
