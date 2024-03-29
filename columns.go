package datev

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"strings"
	"time"
)

func toGermanFloat(s string) string {
	return strings.Replace(s, ".", ",", 1)
}

// 1 - Umsatz (ohne Soll/Haben-Kz)
type amount struct {
	value float64
}

func (d amount) index() int {
	return 1
}

func (d amount) validate() error {
	if d.value == 0 {
		return errors.New("0.00 is not allowed as amount")
	}
	return nil
}

func (d amount) convert() string {
	return toGermanFloat(fmt.Sprintf("%.2f", d.value))
}

// 2 - Soll/Haben-Kennzeichen
type sollHaben struct {
	value string
}

func (s sollHaben) index() int {
	return 2
}

func (s sollHaben) validate() error {
	if s.value == "S" || s.value == "H" {
		return nil
	}
	return fmt.Errorf("%s isn't allowed as Soll/Haben Kennzeichen", s.value)
}

func (s sollHaben) convert() string {
	return s.value
}

// 3 - currency -> WKZ Umsatz
type currency struct {
	value string
}

func (c currency) index() int {
	return 3
}

func (c currency) validate() error {
	if len(c.value) != 3 {
		return errors.New("currency must have a length of 3")
	}
	return nil
}

func (c currency) convert() string {
	return c.value
}

// 4 - course -> Kurs
type course struct {
	value float64
}

func (c course) index() int {
	return 4
}

func (c course) validate() error {
	if c.value <= 0 {
		return errors.New("course must over 0.0")
	}
	return nil
}

func (c course) convert() string {
	return strings.TrimRight(toGermanFloat(fmt.Sprintf("%.6f", c.value)), "0")
}

// 7 - account -> Konto
type account struct {
	value int
}

func (a account) index() int {
	return 7
}

func (a account) validate() error {
	return nil
}

func (a account) convert() string {
	return fmt.Sprintf("%d", a.value)
}

// 8 - cAccount -> Gegenkonto (ohne BU-Schlüssel)
type cAccount struct {
	value int
}

func (a cAccount) index() int {
	return 8
}

func (a cAccount) validate() error {
	return nil
}

func (a cAccount) convert() string {
	return fmt.Sprintf("%d", a.value)
}

// 9 - buKey -> BU-Schlüssel
type buKey struct {
	value int
}

func (b buKey) index() int {
	return 9
}

func (b buKey) validate() error {
	return nil
}

func (b buKey) convert() string {
	return fmt.Sprintf("%d", b.value)
}

// 10 - date -> Belegdatum
type date struct {
	value time.Time
}

func (d date) index() int {
	return 10
}

func (d date) validate() error {
	if reflect.DeepEqual(d.value, time.Time{}) {
		return errors.New("empty time type given")
	}
	return nil
}
func (d date) convert() string {
	return fmt.Sprintf("%d%02d", d.value.Day(), int(d.value.Month()))
}

// 11 - docField -> Belegfeld 1
type docField struct {
	value string
}

func (d docField) index() int {
	return 11
}

func (d docField) validate() error {
	return nil
}

func (d docField) convert() string {
	return d.value
}

// 14 - text -> Buchungstext
type text struct {
	value string
}

func (t text) index() int {
	return 14
}

func (t text) validate() error {
	return nil
}

func (t text) convert() string {
	return t.value
}

// 20 - documentLink -> Beleglink
type documentLink struct {
	value uuid.UUID
}

func (l documentLink) index() int {
	return 20
}

func (l documentLink) validate() error {
	return nil
}

func (l documentLink) convert() string {
	return fmt.Sprintf(`BEDI "%s"`, l.value)
}

// 37 - kost -> KOST1 - Kostenstelle
type kost struct {
	value int
}

func (k kost) index() int {
	return 37
}

func (k kost) validate() error {
	return nil
}

func (k kost) convert() string {
	return fmt.Sprintf("%d", k.value)
}

// 40 - destinationVatIDOrCountry -> EU-Land u. UStID (Bestimmung)
type destinationVatIDOrCountry struct {
	value string
}

func (v destinationVatIDOrCountry) index() int {
	return 40
}

func (v destinationVatIDOrCountry) validate() error {
	return validateVatIDOrCountry(v.value)
}

func (v destinationVatIDOrCountry) convert() string {
	if v.value == "GR" {
		return "EL"
	}
	return v.value
}

// 41 - destinationVatRate -> EU-Steuersatz (Bestimmung)
type destinationVatRate struct {
	value float64
}

func (o destinationVatRate) index() int {
	return 41
}

func (o destinationVatRate) validate() error {
	return validateVatRate(o.value)
}

func (o destinationVatRate) convert() string {
	return toGermanFloat(fmt.Sprintf("%.2f", o.value))
}

// 48 - additionalInformationType1 -> Zusatzinformation – Art 1
type additionalInformationType1 struct {
	value string
}

func (o additionalInformationType1) index() int {
	return 48
}

func (o additionalInformationType1) validate() error {
	return nil
}

func (o additionalInformationType1) convert() string {
	return o.value
}

// 49 - additionalInformationValue1 -> Zusatzinformation – Inhalt 1
type additionalInformationValue1 struct {
	value string
}

func (o additionalInformationValue1) index() int {
	return 49
}

func (o additionalInformationValue1) validate() error {
	return nil
}

func (o additionalInformationValue1) convert() string {
	return o.value
}

// 115 - Leistungsdatum
type performanceDate struct {
	value time.Time
}

func (o performanceDate) index() int {
	return 115
}

func (o performanceDate) validate() error {
	if reflect.DeepEqual(o.value, time.Time{}) {
		return errors.New("Leistungsdatum darf nicht leer sein")
	}
	return nil
}

func (o performanceDate) convert() string {
	return o.value.Format("2012006")
}

// 123 - originVatIDOrCountry -> EU-Land u. USt-IdNr. (Ursprung)
type originVatIDOrCountry struct {
	value string
}

func (v originVatIDOrCountry) index() int {
	return 123
}

func (v originVatIDOrCountry) validate() error {
	return validateVatIDOrCountry(v.value)
}

func (v originVatIDOrCountry) convert() string {
	if v.value == "GR" {
		return "EL"
	}
	return v.value
}

// 124 - originVatRate -> EU-Steuersatz (Ursprung)
type originVatRate struct {
	value float64
}

func (o originVatRate) index() int {
	return 124
}

func (o originVatRate) validate() error {
	return validateVatRate(o.value)
}

func (o originVatRate) convert() string {
	return toGermanFloat(fmt.Sprintf("%.2f", o.value))
}
