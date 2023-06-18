package datev

import (
	"fmt"
	"reflect"
	"time"
)

type bookingColumn interface {
	index() int
	validate() error
	convert() string
}

// errorMap stores errors that occurs while settings values with datev column index as key
type errorMap map[int]string

func (em errorMap) errors() []string {
	s := make([]string, len(em))
	i := 0
	for k, v := range em {
		s[i] = fmt.Sprintf("%s -> %s", columnNames[k-1], v)
		i += 1
	}

	return s
}

func (em errorMap) HasErrors() bool {
	return len(em) != 0
}

type Booking struct {
	values []bookingColumn
	errs   errorMap
	Period Period
}

func (b *Booking) setPeriod(t time.Time) {
	b.Period = Period{
		Month: int(t.Month()),
		Year:  t.Year(),
	}
}

func (b *Booking) String() string {
	s := ""
	for _, col := range b.values {
		if col != nil {
			s += fmt.Sprintf("%3d - %-40s %s -> %s\n", col.index(), columnNames[col.index()-1], "", col.convert())
		}
	}

	return s
}

func (b *Booking) exportValues() []string {
	output := make([]string, len(columnNames))
	for i, col := range b.values {
		if col != nil {
			output[i] = col.convert()
		}
	}

	return output
}

func newBooking() *Booking {
	return &Booking{
		values: make([]bookingColumn, len(columnNames)),
		errs:   make(errorMap, 0),
	}
}

func (b *Booking) setValue(data bookingColumn) {
	// if data isn't valid, we store this information into an errorMap
	if err := data.validate(); err != nil {
		b.errs[data.index()] = err.Error()
	}

	b.values[data.index()-1] = data // index - 1, because we create []bookingColumns with len of columns in DATEV-Format
}

func (b *Booking) IsEqual(other Booking) bool {
	if b.String() != other.String() {
		return false
	}

	if b.Period.String() != other.Period.String() {
		return false
	}

	return true
}

func (b *Booking) IsEmpty() bool {
	return reflect.DeepEqual(*b, Booking{})
}
