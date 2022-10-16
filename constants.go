package datev

const (
	BuKeyNotTaxable    int = 191
	BuKeyReverseCharge int = 270
	BuKeyOSS           int = 280
	BuKeyUSt7          int = 102
	BuKeyUSt19         int = 101
)

func BuKeyByVatRate(rate float64) int {
	if rate > 1 {
		rate /= 100
	}

	switch rate {
	case 0.07:
		return 2
	case 0.19:
		return 3
	case 0.05:
		return 4
	case 0.16:
		return 5
	}
	return 0
}

var columnNames = []string{
	"Umsatz (ohne Soll/Haben-Kz)",     // 1
	"Soll/Haben-Kennzeichen",          // 2
	"WKZ Umsatz",                      // 3
	"Kurs",                            // 4
	"Basis-Umsatz",                    // 5
	"WKZ Basisumsatz",                 // 6
	"Konto",                           // 7
	"Gegenkonto (ohne BU-Schlüssel)",  // 8
	"BU-Schlüssel",                    // 9
	"Belegdatum",                      // 10
	"Belegfeld 1",                     // 11
	"Belegfeld 2",                     // 12
	"Skonto",                          // 13
	"Buchungstext",                    // 14
	"Postensperre",                    // 15
	"Diverse Adressnummer",            // 16
	"Geschäftspartnerbank",            // 17
	"Sachverhalt",                     // 18
	"Zinssperre",                      // 19
	"Beleglink",                       // 20
	"Beleginfo – Art 1",               // 21
	"Beleginfo – Inhalt 1",            // 22
	"Beleginfo – Art 2",               // 23
	"Beleginfo – Inhalt 2",            // 24
	"Beleginfo – Art 3",               // 25
	"Beleginfo – Inhalt 3",            // 26
	"Beleginfo – Art 4",               // 27
	"Beleginfo – Inhalt 4",            // 28
	"Beleginfo – Art 5",               // 29
	"Beleginfo – Inhalt 5",            // 30
	"Beleginfo – Art 6",               // 31
	"Beleginfo – Inhalt 6",            // 32
	"Beleginfo – Art 7",               // 33
	"Beleginfo – Inhalt 7",            // 34
	"Beleginfo – Art 8",               // 35
	"Beleginfo – Inhalt 8",            // 36
	"KOST1 - Kostenstelle",            // 37
	"KOST2 - Kostenstelle",            // 38
	"Kost-Menge",                      // 39
	"EU-Land u. UStID (Bestimmung)",   // 40
	"EU-Steuersatz (Bestimmung)",      // 41
	"Abw. Versteuerungsart",           // 42
	"Sachverhalt L+L",                 // 43
	"Funktionsergänzung L+L",          // 44
	"BU 49 Hauptfunktionstyp",         // 45
	"BU 49 Hauptfunktionsnummer",      // 46
	"BU 49 Funktionsergänzung",        // 47
	"Zusatzinformation – Art 1",       // 48
	"Zusatzinformation – Inhalt 1",    // 49
	"Zusatzinformation – Art 2",       // 50
	"Zusatzinformation – Inhalt 2",    // 51
	"Zusatzinformation – Art 3",       // 52
	"Zusatzinformation – Inhalt 3",    // 53
	"Zusatzinformation – Art 4",       // 54
	"Zusatzinformation – Inhalt 4",    // 55
	"Zusatzinformation – Art 5",       // 56
	"Zusatzinformation – Inhalt 5",    // 57
	"Zusatzinformation – Art 6",       // 58
	"Zusatzinformation – Inhalt 6",    // 59
	"Zusatzinformation – Art 7",       // 60
	"Zusatzinformation – Inhalt 7",    // 61
	"Zusatzinformation – Art 8",       // 62
	"Zusatzinformation – Inhalt 8",    // 63
	"Zusatzinformation – Art 9",       // 64
	"Zusatzinformation – Inhalt 9",    // 65
	"Zusatzinformation – Art 10",      // 66
	"Zusatzinformation – Inhalt 10",   // 67
	"Zusatzinformation – Art 11",      // 68
	"Zusatzinformation – Inhalt 11",   // 69
	"Zusatzinformation – Art 12",      // 70
	"Zusatzinformation – Inhalt 12",   // 71
	"Zusatzinformation – Art 13",      // 72
	"Zusatzinformation – Inhalt 13",   // 73
	"Zusatzinformation – Art 14",      // 74
	"Zusatzinformation – Inhalt 14",   // 75
	"Zusatzinformation – Art 15",      // 76
	"Zusatzinformation – Inhalt 15",   // 77
	"Zusatzinformation – Art 16",      // 78
	"Zusatzinformation – Inhalt 16",   // 79
	"Zusatzinformation – Art 17",      // 80
	"Zusatzinformation – Inhalt 17",   // 81
	"Zusatzinformation – Art 18",      // 82
	"Zusatzinformation – Inhalt 18",   // 83
	"Zusatzinformation – Art 19",      // 84
	"Zusatzinformation – Inhalt 19",   // 85
	"Zusatzinformation – Art 20",      // 86
	"Zusatzinformation – Inhalt 20",   // 87
	"Stück",                           // 88
	"Gewicht",                         // 89
	"Zahlweise",                       // 90
	"Forderungsart",                   // 91
	"Veranlagungsjahr",                // 92
	"Zugeordnete Fälligkeit",          // 93
	"Skontotyp",                       // 94
	"Auftragsnummer",                  // 95
	"Buchungstyp",                     // 96
	"USt-Schlüssel (Anzahlungen)",     // 97
	"EU-Mitgliedstaat (Anzahlungen)",  // 98
	"Sachverhalt L+L (Anzahlungen)",   // 99
	"EU-Steuersatz (Anzahlungen)",     // 100
	"Erlöskonto (Anzahlungen)",        // 101
	"Herkunft-Kz",                     // 102
	"Leerfeld",                        // 103
	"KOST-Datum",                      // 104
	"SEPA-Mandatsreferenz",            // 105
	"Skontosperre",                    // 106
	"Gesellschaftername",              // 107
	"Beteiligtennummer",               // 108
	"Identifikationsnummer",           // 109
	"Zeichnernummer",                  // 110
	"Postensperre bis",                // 111
	"Bezeichnung SoBil-Sachverhalt",   // 112
	"Kennzeichen SoBil-Buchung",       // 113
	"Festschreibung",                  // 114
	"Leistungsdatum",                  // 115
	"Datum Zuord. Steuerperiode",      // 116
	"Fälligkeit",                      // 117
	"Generalumkehr",                   // 118
	"Steuersatz",                      // 119
	"Land",                            // 120
	"Abrechnungsreferenz",             // 121
	"BVV-Position",                    // 122
	"EU-Land u. USt-IdNr. (Ursprung)", // 123
	"EU-Steuersatz (Ursprung)",        // 124
}
