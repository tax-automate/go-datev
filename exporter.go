package datev

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"path/filepath"
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
	ConsultantNumber       int
	ClientNumber           int
	SKL                    int
	SKR                    int
	Fixation               _bool
	SplitBookingsByDebitor bool
}

type Exporter struct {
	filePath    string
	cfg         ExporterConfig
	financeYear time.Time
	period      Period
}

func NewExporter(filePath string, cfg ExporterConfig, period Period) Exporter {
	financeYear := time.Date(period.Begin.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	return Exporter{filePath: filePath, cfg: cfg, period: period, financeYear: financeYear}
}

func (e Exporter) SetDeviatingFinanceYear(year, month int) {
	e.financeYear = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
}

type Period struct {
	Begin time.Time
	End   time.Time
}

func (e Exporter) CreateExport(bookings []Booking, fileName string) error {
	if !e.cfg.SplitBookingsByDebitor {
		return e.CreateExportFile(bookings, fileName)
	}

	debitorBookings := make([]Booking, 0)
	otherBookings := make([]Booking, 0)
	minDebitorAcc := int(math.Pow(10, float64(e.cfg.SKL)))
	for _, booking := range bookings {
		acc := booking.values[6]._value().(int)
		cAcc := booking.values[7]._value().(int)

		if acc >= minDebitorAcc || cAcc >= minDebitorAcc {
			debitorBookings = append(debitorBookings, booking)
		} else {
			otherBookings = append(otherBookings, booking)
		}
	}

	err := e.CreateExportFile(debitorBookings, fileName+" Debitoren")
	if err != nil {
		return err
	}

	err = e.CreateExportFile(otherBookings, fileName)
	if err != nil {
		return err
	}
	return nil
}

func (e Exporter) CreateExportFile(bookings []Booking, fileName string) error {
	fileName = fmt.Sprintf("EXTF_%s.csv", fileName)
	path := filepath.Join(e.filePath, fileName)

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	bomUtf8 := []byte{0xEF, 0xBB, 0xBF}
	_, err = f.Write(bomUtf8)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(f)
	defer func() {
		writer.Flush()
		_ = f.Close()
	}()
	writer.Comma = ';'

	// Header
	err = writer.Write(e.createHeaderRow(fileName))
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
		fmt.Sprintf("%d", e.cfg.ConsultantNumber),                                               // Beraternummer
		fmt.Sprintf("%d", e.cfg.ClientNumber),                                                   // Mandantennummer
		fmt.Sprintf("%d%d%d", e.financeYear.Year(), e.financeYear.Month(), e.financeYear.Day()), // Finanzjahr
		fmt.Sprintf("%d", e.cfg.SKL),                                                            // SKL
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
