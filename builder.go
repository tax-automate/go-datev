package datev

import (
	"math"
	"time"
)

type BookingBuilder struct {
	b   *Booking
	lgr *BookingLogger
	cnt int
}

func NewBookingBuilder() *BookingBuilder {
	return &BookingBuilder{b: newBooking(), lgr: NewBookingLogger()}
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
		// after returning the booking, set booking and errors to empty values
		bb.b = newBooking()
		bb.cnt += 1
	}()

	if bb.b.errs.HasErrors() {
		bb.lgr.addError(bb.cnt, bb.b.errs.errors())
	}

	return *bb.b
}

func (bb *BookingBuilder) setValue(data bookingColumn) *BookingBuilder {
	bb.b.setValue(data)
	return bb
}

func (bb *BookingBuilder) SetAmount(n float64) *BookingBuilder {
	if n > 0 {
		bb.b.setValue(sollHaben{"S"})
	} else {
		bb.b.setValue(sollHaben{"H"})
	}
	return bb.setValue(amount{math.Abs(n)})
}

func (bb *BookingBuilder) SetCurrency(curr string, _course float64) *BookingBuilder {
	bb.b.setValue(currency{curr})
	bb.b.setValue(course{_course})
	return bb
}

func (bb *BookingBuilder) SetAccount(n int) *BookingBuilder {
	return bb.setValue(account{n})
}

func (bb *BookingBuilder) SetCAccount(n int) *BookingBuilder {
	return bb.setValue(cAccount{n})
}

func (bb *BookingBuilder) SetBuKey(n int) *BookingBuilder {
	return bb.setValue(buKey{n})
}

func (bb *BookingBuilder) SetDate(t time.Time) *BookingBuilder {
	return bb.setValue(date{value: t})
}

func (bb *BookingBuilder) SetDocField(s string) *BookingBuilder {
	return bb.setValue(docField{s})
}

func (bb *BookingBuilder) SetText(s string) *BookingBuilder {
	return bb.setValue(text{s})
}

func (bb *BookingBuilder) SetKOST(n int) *BookingBuilder {
	return bb.setValue(kost{n})
}

func (bb *BookingBuilder) SetVatID(s string) *BookingBuilder {
	return bb.setValue(destinationVatIDOrCountry{s})
}

func (bb *BookingBuilder) SetDestinationEuInformation(countryCode string, rate float64) *BookingBuilder {
	bb.b.setValue(destinationVatIDOrCountry{countryCode})
	bb.b.setValue(destinationVatRate{rate})
	return bb
}

func (bb *BookingBuilder) SetOriginEuInformation(countryCode string, rate float64) *BookingBuilder {
	bb.b.setValue(originVatIDOrCountry{countryCode})
	bb.b.setValue(originVatRate{rate})
	return bb
}
