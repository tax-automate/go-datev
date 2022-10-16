package datev

import (
	"regexp"
	"strings"
)

// IsValidVatID check if given string is a valid VatID.
// see https://de.wikipedia.org/wiki/Umsatzsteuer-Identifikationsnummer
func IsValidVatID(vatID string) bool {
	countries := []string{
		"ATU[0-9]{8}",
		"BE[01][0-9]{9}",
		"BG[0-9]{9,10}",
		"HR[0-9]{11}",
		"CY[A-Z0-9]{9}",
		"CZ[0-9]{8,10}",
		"DK[0-9]{8}",
		"EE[0-9]{9}",
		"FI[0-9]{8}",
		"FR[0-9A-Z]{2}[0-9]{9}",
		"DE[0-9]{9}",
		"EL[0-9]{9}",
		"HU[0-9]{8}",
		"IE([0-9]{7}[A-Z]{1,2}|[0-9][A-Z][0-9]{5}[A-Z])",
		"IT[0-9]{11}",
		"LV[0-9]{11}",
		"LT([0-9]{9}|[0-9]{12})",
		"LU[0-9]{8}",
		"MT[0-9]{8}",
		"NL[0-9]{9}B[0-9]{2}",
		"PL[0-9]{10}",
		"PT[0-9]{9}",
		"RO[0-9]{2,10}",
		"SK[0-9]{10}",
		"SI[0-9]{8}",
		"ES[A-Z]([0-9]{8}|[0-9]{7}[A-Z])",
		"SE[0-9]{12}",
		"XI[0-9]{9, 12}",
	}
	pattern := regexp.MustCompile(strings.Join(countries, "|"))

	return pattern.MatchString(vatID)
}
