package datev

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

// 1 - Umsatz (ohne Soll/Haben-Kz)
type amount struct {
	value float64
}

func (d amount) index() int {
	return 1
}

func (d amount) validate(value interface{}) error {
	if value.(float64) == 0 {
		return errors.New("0.00 is not allowed as amount")
	}
	return nil
}

func (d amount) convert() string {
	return fmt.Sprintf("%.2f", d.value)
}

// 2 - Soll/Haben-Kennzeichen
type sollHaben struct {
	value string
}

func (s sollHaben) index() int {
	return 2
}

func (s sollHaben) validate(value interface{}) error {
	v := value.(string)
	if v == "S" || v == "H" {
		return nil
	}
	return fmt.Errorf("%s isn't allowed as Soll/Haben Kennzeichen", v)
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

func (c currency) validate(value interface{}) error {
	if len(value.(string)) != 3 {
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

func (c course) validate(value interface{}) error {
	if value.(float64) <= 0 {
		return errors.New("course must over 0.0")
	}
	return nil
}

func (c course) convert() string {
	return fmt.Sprintf("%.6f", c.value)
}

// 7 - account -> Konto
type account struct {
	value int
}

func (a account) index() int {
	return 7
}

func (a account) validate(value interface{}) error {
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

func (a cAccount) validate(_ interface{}) error {
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

func (b buKey) validate(_ interface{}) error {
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

func (d date) validate(value interface{}) error {
	date := value.(time.Time)
	if !reflect.DeepEqual(date, time.Time{}) {
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

func (d docField) validate(_ interface{}) error {
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

func (t text) validate(_ interface{}) error {
	return nil
}

func (t text) convert() string {
	return t.value
}

// 37 - kost -> KOST1 - Kostenstelle
type kost struct {
	value int
}

func (k kost) index() int {
	return 37
}

func (k kost) validate(_ interface{}) error {
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

func (v destinationVatIDOrCountry) validate(value interface{}) error {
	vatID := value.(string)
	if !IsValidVatID(vatID) {
		return fmt.Errorf("given vatID %s is not valid", vatID)
	}
	return nil
}

func (v destinationVatIDOrCountry) convert() string {
	return v.value
}

// 41 - destinationVatRate -> EU-Steuersatz (Bestimmung)
type destinationVatRate struct {
	value float64
}

func (o destinationVatRate) index() int {
	return 41
}

func (o destinationVatRate) validate(value interface{}) error {
	// TODO: check if given country code is given
	if value.(float64) <= 0 {
		return errors.New("vat rate must be over 0")
	}
	return nil
}

func (o destinationVatRate) convert() string {
	return fmt.Sprintf("%.2f", o.value)
}

// 123 - originVatIDOrCountry -> EU-Land u. USt-IdNr. (Ursprung)
type originVatIDOrCountry struct {
	value string
}

func (v originVatIDOrCountry) index() int {
	return 123
}

func (v originVatIDOrCountry) validate(value interface{}) error {
	vatID := value.(string)
	if !IsValidVatID(vatID) {
		return fmt.Errorf("given vatID %s is not valid", vatID)
	}
	return nil
}

func (v originVatIDOrCountry) convert() string {
	return v.value
}

// 124 - originVatRate -> EU-Steuersatz (Ursprung)
type originVatRate struct {
	value float64
}

func (o originVatRate) index() int {
	return 124
}

func (o originVatRate) validate(value interface{}) error {
	// TODO: check if given country code is given
	if value.(float64) <= 0 {
		return errors.New("vat rate must be over 0")
	}
	return nil
}

func (o originVatRate) convert() string {
	return fmt.Sprintf("%.2f", o.value)
}
