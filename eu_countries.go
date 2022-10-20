package datev

import (
	"reflect"
	"time"
)

type Countries map[string]country

var countriesFrom = time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)

func getEUCountries() Countries {
	output := make(Countries, 0)
	for k, v := range euCountries {
		if !reflect.DeepEqual(v.LeftEUAt, time.Time{}) {
			if v.LeftEUAt.Before(countriesFrom) || v.LeftEUAt.Equal(countriesFrom) {
				continue
			}
		}
		output[k] = v
	}

	return output
}

type country struct {
	InEUSince time.Time
	LeftEUAt  time.Time
}

var euCountries = Countries{
	"AT": country{
		InEUSince: time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	"BE": country{
		InEUSince: time.Date(1958, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	"BG": country{
		InEUSince: time.Date(2007, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	"HR": country{
		InEUSince: time.Date(2013, 7, 1, 0, 0, 0, 0, time.UTC),
	},
	"CY": country{
		InEUSince: time.Date(2004, 5, 1, 0, 0, 0, 0, time.UTC),
	},
	"CZ": country{
		InEUSince: time.Date(2004, 5, 1, 0, 0, 0, 0, time.UTC),
	},
	"DK": country{
		InEUSince: time.Date(1973, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	"EE": country{
		InEUSince: time.Date(2004, 5, 1, 0, 0, 0, 0, time.UTC),
	},
	"FI": country{
		InEUSince: time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	"FR": country{
		InEUSince: time.Date(1958, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	"GR": country{
		InEUSince: time.Date(1981, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	"HU": country{
		InEUSince: time.Date(2004, 5, 1, 0, 0, 0, 0, time.UTC),
	},
	"IE": country{
		InEUSince: time.Date(1973, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	"IT": country{
		InEUSince: time.Date(1958, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	"LV": country{
		InEUSince: time.Date(2004, 5, 1, 0, 0, 0, 0, time.UTC),
	},
	"LT": country{
		InEUSince: time.Date(2004, 5, 1, 0, 0, 0, 0, time.UTC),
	},
	"LU": country{
		InEUSince: time.Date(1958, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	"MT": country{
		InEUSince: time.Date(2004, 5, 1, 0, 0, 0, 0, time.UTC),
	},
	"NL": country{
		InEUSince: time.Date(1958, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	"PL": country{
		InEUSince: time.Date(2004, 5, 1, 0, 0, 0, 0, time.UTC),
	},
	"PT": country{
		InEUSince: time.Date(1986, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	"RO": country{
		InEUSince: time.Date(2007, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	"SK": country{
		InEUSince: time.Date(2004, 5, 1, 0, 0, 0, 0, time.UTC),
	},
	"SI": country{
		InEUSince: time.Date(2004, 5, 1, 0, 0, 0, 0, time.UTC),
	},
	"ES": country{
		InEUSince: time.Date(1986, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	"SE": country{
		InEUSince: time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	"GB": country{
		InEUSince: time.Date(1973, 1, 1, 0, 0, 0, 0, time.UTC),
		LeftEUAt:  time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC),
	},
	"XI": country{
		InEUSince: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	},
}
