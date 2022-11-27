package datev

import (
	"testing"
	"time"
)

func TestMultipleColumns(t *testing.T) {
	tests := []struct {
		name                  string
		col                   bookingColumn
		isValid               bool
		wantedConvertedString string
	}{
		{
			name:                  "amount with valid",
			col:                   amount{42.9993},
			isValid:               true,
			wantedConvertedString: "43,00",
		},
		{
			name:                  "negative amount",
			col:                   amount{-0.0133},
			isValid:               true,
			wantedConvertedString: "-0,01",
		},
		{
			name:                  "invalid currency",
			col:                   currency{"BLUB"},
			isValid:               false,
			wantedConvertedString: "",
		},
		{
			name:                  "valid currency",
			col:                   currency{"USD"},
			isValid:               true,
			wantedConvertedString: "USD",
		},
		{
			name:                  "valid course",
			col:                   course{1.2546},
			isValid:               true,
			wantedConvertedString: "1,2546",
		},
		{
			name:                  "valid course with more than 6 decimals",
			col:                   course{9.25467423},
			isValid:               true,
			wantedConvertedString: "9,254674",
		},
		{
			name:                  "date",
			col:                   date{time.Date(2022, 11, 30, 0, 0, 0, 0, time.UTC)},
			isValid:               true,
			wantedConvertedString: "3011",
		},
		{
			name:                  "empty date",
			col:                   date{time.Time{}},
			isValid:               false,
			wantedConvertedString: "",
		},
		{
			name:                  "valid vat id",
			col:                   destinationVatIDOrCountry{"DE999999999"},
			isValid:               true,
			wantedConvertedString: "DE999999999",
		},
		{
			name:                  "invalid vat id",
			col:                   destinationVatIDOrCountry{"DE9999dsadsa99999"},
			isValid:               false,
			wantedConvertedString: "",
		},
		{
			name:                  "valid country",
			col:                   destinationVatIDOrCountry{"PL"},
			isValid:               true,
			wantedConvertedString: "PL",
		},
		{
			name:                  "invalid country",
			col:                   originVatIDOrCountry{"US"},
			isValid:               false,
			wantedConvertedString: "",
		},
		{
			name:                  "special case GR -> EL",
			col:                   destinationVatIDOrCountry{"GR"},
			isValid:               true,
			wantedConvertedString: "EL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.col.validate()
			if err != nil && tt.isValid {
				t.Fatalf("unwanted errors -> %v", err)
			}

			if got := tt.col.convert(); got != tt.wantedConvertedString && tt.isValid {
				t.Errorf("converted strings: got %q, want %q", got, tt.wantedConvertedString)
			}
		})
	}
}
