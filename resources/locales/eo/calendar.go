package eo

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d-'a' 'de' MMMM y", Long: "y-MMMM-dd", Medium: "y-MMM-dd", Short: "yy-MM-dd"},
			Time:     ut.CalendarDateFormat{Full: "H-'a' 'horo' 'kaj' m:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "jan", Feb: "feb", Mar: "mar", Apr: "apr", May: "maj", Jun: "jun", Jul: "jul", Aug: "aŭg", Sep: "sep", Oct: "okt", Nov: "nov", Dec: "dec"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "januaro", Feb: "februaro", Mar: "marto", Apr: "aprilo", May: "majo", Jun: "junio", Jul: "julio", Aug: "aŭgusto", Sep: "septembro", Oct: "oktobro", Nov: "novembro", Dec: "decembro"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "di", Mon: "lu", Tue: "ma", Wed: "me", Thu: "ĵa", Fri: "ve", Sat: "sa"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "M", Thu: "Ĵ", Fri: "V", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "dimanĉo", Mon: "lundo", Tue: "mardo", Wed: "merkredo", Thu: "ĵaŭdo", Fri: "vendredo", Sat: "sabato"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "atm", PM: "ptm"},
			},
		},
	}
}
