package mfe

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "zan", Feb: "fev", Mar: "mar", Apr: "avr", May: "me", Jun: "zin", Jul: "zil", Aug: "out", Sep: "sep", Oct: "okt", Nov: "nov", Dec: "des"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "z", Feb: "f", Mar: "m", Apr: "a", May: "m", Jun: "z", Jul: "z", Aug: "o", Sep: "s", Oct: "o", Nov: "n", Dec: "d"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "zanvie", Feb: "fevriye", Mar: "mars", Apr: "avril", May: "me", Jun: "zin", Jul: "zilye", Aug: "out", Sep: "septam", Oct: "oktob", Nov: "novam", Dec: "desam"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "dim", Mon: "lin", Tue: "mar", Wed: "mer", Thu: "ze", Fri: "van", Sat: "sam"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "d", Mon: "l", Tue: "m", Wed: "m", Thu: "z", Fri: "v", Sat: "s"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "dimans", Mon: "lindi", Tue: "mardi", Wed: "merkredi", Thu: "zedi", Fri: "vandredi", Sat: "samdi"},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
