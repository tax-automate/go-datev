package datev

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	personCode   = "TA"
	exportedFrom = "tax-automate"
)

type ExporterConfig struct {
	ConsultantNumber       int
	ClientNumber           int
	SKL                    int
	SKR                    int
	Fixation               bool
	SplitBookingsByDebitor bool
}

type Exporter struct {
	filePath    string
	cfg         ExporterConfig
	financeYear time.Time
	period      Period
}

func NewExporter(filePath string, cfg ExporterConfig) *Exporter {
	return &Exporter{filePath: filePath, cfg: cfg}
}

func (e *Exporter) SetDeviatingFinanceYear(year, month int) {
	e.financeYear = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
}

type Period struct {
	month int
	year  int
}

func (p Period) String() string {
	return fmt.Sprintf("%02d-%d", p.month, p.year)
}

func (p Period) Begin() time.Time {
	return time.Date(p.year, time.Month(p.month), 1, 0, 0, 0, 0, time.UTC)
}

func (p Period) End() time.Time {
	return time.Date(p.year, time.Month(p.month), 1, 0, 0, 0, 0, time.UTC).AddDate(0, 1, -1)
}

// CreateExport creates an export file in DATEV-Format
func (e *Exporter) CreateExport(bookings []Booking, fileName string) error {
	err := os.MkdirAll(e.filePath, os.ModePerm)
	if err != nil {
		return err
	}

	sortedBookings := sortBookingsByPeriod(bookings)
	mainPeriod := getMainPeriod(sortedBookings)
	for period, bookingsForFile := range sortedBookings {
		e.period = period
		e.financeYear = time.Date(period.year, time.Month(period.month), 1, 0, 0, 0, 0, time.UTC)

		if mainPeriod != period {
			fileName = fmt.Sprintf("%s - Zeitraum %s", fileName, period.String())
		}
		err = e.writeFile(bookingsForFile, fileName)
		if err != nil {
			return err
		}
	}

	return nil
}

// getMainPeriod looks what period have the most bookings. This Period will be the main period
func getMainPeriod(data map[Period][]Booking) Period {
	var p Period
	var maxLength int
	for period, bookings := range data {
		if len(bookings) > maxLength {
			maxLength = len(bookings)
			p = period
		}
	}

	return p
}

func sortBookingsByPeriod(bookings []Booking) map[Period][]Booking {
	output := make(map[Period][]Booking, 0)
	for _, booking := range bookings {
		p := booking.Period
		if _, ok := output[p]; !ok {
			output[p] = []Booking{}
		}
		output[p] = append(output[p], booking)
	}
	return output
}

func (e *Exporter) writeFile(bookings []Booking, fileName string) error {
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
func (e *Exporter) createHeaderRow(fileName string) []string {
	now := time.Now()
	header := []string{
		"EXTF",           // Format
		"700",            // Versions Nr
		"21",             // category
		"Buchungsstapel", // format name
		"12",             // Format version
		fmt.Sprintf("%d%02d%02d%02d%02d%02d000", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()), // created at
		"",           // Imported (must be empty)
		"RE",         // origin
		exportedFrom, // exported from
		"",           // imported from (must be empty)
		fmt.Sprintf("%d", e.cfg.ConsultantNumber),                                                   // Beraternummer
		fmt.Sprintf("%d", e.cfg.ClientNumber),                                                       // Mandantennummer
		fmt.Sprintf("%d%02d%02d", e.financeYear.Year(), e.financeYear.Month(), e.financeYear.Day()), // Finanzjahr
		fmt.Sprintf("%d", e.cfg.SKL),                                                                // SKL
		fmt.Sprintf("%d%02d%02d", e.period.year, e.period.month, e.period.Begin().Day()),
		fmt.Sprintf("%d%02d%02d", e.period.year, e.period.month, e.period.End().Day()),
		fileName,
		personCode,                   // person code
		"1",                          // booking_type  # 1 = Fibu / 2 = Jahresabschluss
		"",                           // accounting_purpose
		boolAsString(e.cfg.Fixation), // fixation  # Festschreibung 1 = Ja / 0 = Nein
		"EUR",                        // Currency
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

func boolAsString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}
