package datev

import (
	"math"
	"time"
)

type BookingBuilder struct {
	b *Booking
}

func NewBookingBuilder() *BookingBuilder {
	return &BookingBuilder{b: newBooking()}
}

// SetMinValues is a helper function to set minimal attributes to Booking
func (bb *BookingBuilder) SetMinValues(date time.Time, amount float64, cAccount int, account int, docField string, text string) *BookingBuilder {
	bb.SetDate(date)
	bb.SetAmount(amount)
	bb.SetCAccount(cAccount)
	bb.SetAccount(account)
	bb.SetDocField(docField)
	bb.SetText(text)

	return bb
}

func (bb *BookingBuilder) Build() Booking {
	defer func() {
		// after returning the booking, set booking to empty values
		bb.b = newBooking()
	}()
	return *bb.b
}

func (bb *BookingBuilder) setValue(data bookingColumn) *BookingBuilder {
	bb.b.setValue(data)
	return bb
}

func (bb *BookingBuilder) SetDate(t time.Time) *BookingBuilder {
	return bb.setValue(date{value: t})
}

func (bb *BookingBuilder) SetAmount(n float64) *BookingBuilder {
	if n > 0 {
		bb.b.setValue(sollHaben{"S"})
	} else {
		bb.b.setValue(sollHaben{"H"})
	}
	return bb.setValue(amount{value: math.Abs(n)})
}

func (bb *BookingBuilder) SetCAccount(n int) *BookingBuilder {
	panic("Not implemented yet!")
}

func (bb *BookingBuilder) SetAccount(n int) *BookingBuilder {
	panic("Not implemented yet!")
}

func (bb *BookingBuilder) SetText(s string) *BookingBuilder {
	panic("Not implemented yet!")
}

func (bb *BookingBuilder) SetCurrency(curr string, course float64) *BookingBuilder {
	panic("Not implemented yet!")
}

func (bb *BookingBuilder) SetVatId(s string) *BookingBuilder {
	panic("Not implemented yet!")
}

func (bb *BookingBuilder) SetEuInformation(s string, n float64) *BookingBuilder {
	panic("Not implemented yet!")
}

func (bb *BookingBuilder) SetVatIdOrigin(s string) *BookingBuilder {
	panic("Not implemented yet!")
}

func (bb *BookingBuilder) SetEuInformationOrigin(s string, n float64) *BookingBuilder {
	panic("Not implemented yet!")
}

func (bb *BookingBuilder) SetBuKey(n int) *BookingBuilder {
	panic("Not implemented yet!")
}

func (bb *BookingBuilder) SetDocField(s string) *BookingBuilder {
	panic("Not implemented yet!")
}

func (bb *BookingBuilder) SetKOST(n int) *BookingBuilder {
	panic("Not implemented yet!")
}
