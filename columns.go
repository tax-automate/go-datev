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

// 10 - Belegdatum
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
