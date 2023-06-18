package datev

import (
	"reflect"
	"testing"
	"time"
)

func Test_sortBookingsByPeriod(t *testing.T) {
	builder := NewBookingBuilder()
	tests := []struct {
		name     string
		bookings []Booking
		want     map[Period][]Booking
	}{
		{
			name: "multiple periods",
			bookings: []Booking{
				builder.SetDate(time.Date(2023, 3, 15, 0, 0, 0, 0, time.UTC)).Build(),
				builder.SetDate(time.Date(2023, 4, 15, 0, 0, 0, 0, time.UTC)).Build(),
			},
			want: map[Period][]Booking{
				Period{
					Month: 3,
					Year:  2023,
				}: {builder.SetDate(time.Date(2023, 3, 15, 0, 0, 0, 0, time.UTC)).Build()},
				Period{
					Month: 4,
					Year:  2023,
				}: {builder.SetDate(time.Date(2023, 4, 15, 0, 0, 0, 0, time.UTC)).Build()},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortBookingsByPeriod(tt.bookings); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortBookingsByPeriod() = %v, want %v", got, tt.want)
			}
		})
	}
}
