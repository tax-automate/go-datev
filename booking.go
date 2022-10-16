package datev

type bookingColumn interface {
	index() int
	validate(value interface{}) error
	convert() string
}

type Booking struct {
	values []bookingColumn
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
