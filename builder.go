package datev

import (
	"encoding/csv"
	"math"
	"time"
)

// ErrorMap stores errors that occurs while settings values with datev column index as key
type ErrorMap map[int]string

func (em ErrorMap) ExportToCsv(writer *csv.Writer) {
	for k, v := range em {
		s := []string{columnNames[k-1], v + "\n"}
		_ = writer.Write(s)
	}
}

type BookingBuilder struct {
	b    *Booking
	errs ErrorMap
}

func NewBookingBuilder() *BookingBuilder {
	return &BookingBuilder{b: newBooking(), errs: make(ErrorMap, 0)}
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

func (bb *BookingBuilder) Build() (Booking, ErrorMap) {
	defer func() {
		// after returning the booking, set booking and errors to empty values
		bb.b = newBooking()
		bb.errs = make(ErrorMap, 0)
	}()
	return *bb.b, bb.errs
}

func (bb *BookingBuilder) setValue(data bookingColumn) *BookingBuilder {
	if err := data.validate(); err != nil {
		bb.errs[data.index()] = err.Error()
		return bb
	}
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
	return bb.setValue(originVatIDOrCountry{s})
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
