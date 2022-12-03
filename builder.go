package datev

import (
	"math"
	"time"
)

// BookingBuilder implements a builder pattern to create booking that are equivalent to the format specification from
// the DATEV-Format
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

// Build returns the prepared Booking
// Before returning, the function checks if any errors occurred while preparing the booking
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

// SetAmount sets the values for the datev columns 'Umsatz (ohne Soll/Haben-Kz)' and 'Soll/Haben-Kennzeichen'.
// The 'Soll/Haben-Kennzeichen' is implicit set. 'S' if the amount is positive and 'H' if it's not
func (bb *BookingBuilder) SetAmount(n float64) *BookingBuilder {
	if n > 0 {
		bb.b.setValue(sollHaben{"S"})
	} else {
		bb.b.setValue(sollHaben{"H"})
	}
	return bb.setValue(amount{math.Abs(n)})
}

// SetCurrency sets the values for the datev columns 'WKZ-Umsatz' and 'Kurs'
func (bb *BookingBuilder) SetCurrency(curr string, _course float64) *BookingBuilder {
	bb.b.setValue(currency{curr})
	bb.b.setValue(course{_course})
	return bb
}

// SetAccount set the value for the datev column 'Konto'
func (bb *BookingBuilder) SetAccount(n int) *BookingBuilder {
	return bb.setValue(account{n})
}

// SetCAccount set the value for the datev column 'Gegenkonto'
func (bb *BookingBuilder) SetCAccount(n int) *BookingBuilder {
	return bb.setValue(cAccount{n})
}

// SetBuKey set the value for the datev column 'BU-Schlüssel'
func (bb *BookingBuilder) SetBuKey(n int) *BookingBuilder {
	return bb.setValue(buKey{n})
}

// SetDate set the value for the datev column 'Belegdatum'
func (bb *BookingBuilder) SetDate(t time.Time) *BookingBuilder {
	return bb.setValue(date{value: t})
}

// SetDocField set the value for the datev column 'Belegfeld 1'
func (bb *BookingBuilder) SetDocField(s string) *BookingBuilder {
	return bb.setValue(docField{s})
}

// SetText set the value for the datev column 'Buchungstext'
func (bb *BookingBuilder) SetText(s string) *BookingBuilder {
	return bb.setValue(text{s})
}

// SetKOST set the value for the datev column 'KOST1'
func (bb *BookingBuilder) SetKOST(n int) *BookingBuilder {
	return bb.setValue(kost{n})
}

// SetVatID set the value for the datev column 'EU-Land u. UStID (Bestimmung)'
func (bb *BookingBuilder) SetVatID(s string) *BookingBuilder {
	return bb.setValue(destinationVatIDOrCountry{s})
}

// SetDestinationEuInformation sets the values for the datev columns 'EU-Land u. UStID (Bestimmung)' and 'EU-Steuersatz (Bestimmung)'
// Before saving, the function sanitize the vat rate to integer level, because the DATEV-Formats expect this
func (bb *BookingBuilder) SetDestinationEuInformation(countryCode string, rate float64) *BookingBuilder {
	if rate < 1 {
		rate *= 100
	}
	bb.b.setValue(destinationVatIDOrCountry{countryCode})
	bb.b.setValue(destinationVatRate{rate})
	return bb
}

// SetOriginEuInformation sets the values for the datev columns 'EU-Land u. UStID (Ursprung)' and 'EU-Steuersatz (Ursprung)'
// Before saving, the function sanitize the vat rate to integer level, because the DATEV-Formats expect this
func (bb *BookingBuilder) SetOriginEuInformation(countryCode string, rate float64) *BookingBuilder {
	if rate < 1 {
		rate *= 100
	}
	bb.b.setValue(originVatIDOrCountry{countryCode})
	bb.b.setValue(originVatRate{rate})
	return bb
}
