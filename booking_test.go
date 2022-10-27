package datev

import (
	"reflect"
	"testing"
	"time"
)

func TestBookingToExport(t *testing.T) {
	date := func(year, month, day int) time.Time {
		return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	}

	builder := NewBookingBuilder()
	tests := []struct {
		booking  Booking
		valueMap map[int]string
	}{
		{
			booking: builder.
				SetAmount(-1337.511).
				SetDate(date(2022, 11, 15)).
				SetAccount(1000).
				SetCAccount(9999).
				SetText("Ich bin ein Test").
				SetKOST(1234).
				SetOriginEuInformation("AT", 20.0).
				Build(),
			valueMap: map[int]string{
				1:   "1337,51",
				2:   "H",
				7:   "1000",
				8:   "9999",
				10:  "1511",
				14:  "Ich bin ein Test",
				37:  "1234",
				123: "AT",
				124: "20,00",
			},
		},
	}
	for _, tt := range tests {
		s := make([]string, len(columnNames))
		for index, value := range tt.valueMap {
			s[index-1] = value
		}

		values := tt.booking.exportValues()
		if !reflect.DeepEqual(values, s) {
			t.Errorf("got %v, want %v", values, s)
		}
	}
}
