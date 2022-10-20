package datev

import "fmt"

type bookingColumn interface {
	index() int
	validate() error
	convert() string
}

type Booking struct {
	values []bookingColumn
}

func (b *Booking) String() string {
	s := ""
	for _, col := range b.values {
		s += fmt.Sprintf("%3d - %-40s %s -> %s\n", col.index(), columnNames[col.index()-1], "", col.convert())
	}

	return s
}

func (b *Booking) exportValues() []string {
	output := make([]string, len(columnNames))
	for i, col := range b.values {
		if col != nil {
			output[i] = col.convert()
		} else {
			output[i] = ""
		}
	}

	return output
}

func newBooking() *Booking {
	values := make([]bookingColumn, len(columnNames))
	return &Booking{values: values}
}

func (b *Booking) setValue(value bookingColumn) {
	b.values[value.index()] = value
}
