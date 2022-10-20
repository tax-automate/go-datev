package datev

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

const (
	personCode   = "TA"
	exportedFrom = "tax-automate"
)

type _bool bool

func (b _bool) String() string {
	if b {
		return "1"
	}
	return "0"
}

type ExporterConfig struct {
	ConsultantNumber int
	ClientNumber     int
	FinanceYear      int
	SKL              int
	SKR              int
	Fixation         _bool
}

type Exporter struct {
	filePath string
	cfg      ExporterConfig
	period   Period
}

type Period struct {
	Begin time.Time
	End   time.Time
}

func (e Exporter) CreateExportFile(bookings []Booking, fileName string, cfg ExporterConfig) error {
	writer := csv.NewWriter(os.Stdout)
	writer.Comma = ';'

	// Header
	err := writer.Write(e.createHeaderRow(fileName))
	if err != nil {
		return fmt.Errorf("error while creating header row -> %q", err.Error())
	}

	// column names
	err = writer.Write(columnNames)
	if err != nil {
		return fmt.Errorf("error while creating columns -> %q", err.Error())
	}
	// bookings
	for _, booking := range bookings {
		err = writer.Write(booking.exportValues())
		if err != nil {
			return fmt.Errorf("error %q while creating booking with value %v", err.Error(), booking.String())
		}
	}

	return nil
}

// createHeaderRow see: https://developer.datev.de/datev/platform/de/dtvf/formate/header
func (e Exporter) createHeaderRow(fileName string) []string {
	now := time.Now()
	header := []string{
		"EXTF",           // Format
		"700",            // Versions Nr
		"21",             // category
		"Buchungsstapel", // format name
		"9",              // Format version
		fmt.Sprintf("%d%02d%02d%02d%02d%02d000", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()), // created at
		"",           // Imported (must be empty)
		"RE",         // origin
		exportedFrom, // exported from
		"",           // imported from (must be empty)
		fmt.Sprintf("%d", e.cfg.ConsultantNumber), // Beraternummer
		fmt.Sprintf("%d", e.cfg.ClientNumber),     // Mandantennummer
		fmt.Sprintf("%d", e.cfg.FinanceYear),      // Finanzjahr
		fmt.Sprintf("%d", e.cfg.SKL),              // SKL
		fmt.Sprintf("%d%02d%02d", e.period.Begin.Year(), e.period.Begin.Month(), e.period.Begin.Day()),
		fmt.Sprintf("%d%02d%02d", e.period.End.Year(), e.period.End.Month(), e.period.End.Day()),
		fileName,
		personCode,              // person code
		"1",                     // booking_type  # 1 = Fibu / 2 = Jahresabschluss
		"",                      // accounting_purpose
		e.cfg.Fixation.String(), // fixation  # Festschreibung 1 = Ja / 0 = Nein
		"EUR",                   // Currency
		"",
		"",
		"",
		"",
		fmt.Sprintf("%d", e.cfg.SKR),
		"",
		"",
		"",
		"",
	}

	return header
}
