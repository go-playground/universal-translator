package tk

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "d MMMM y EEEE", Long: "d MMMM y", Medium: "d MMM y", Short: "dd.MM.y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "ýan", Feb: "few", Mar: "mart", Apr: "apr", May: "maý", Jun: "iýun", Jul: "iýul", Aug: "awg", Sep: "sen", Oct: "okt", Nov: "noý", Dec: "dek"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "Ý", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "I", Jul: "I", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "ýanwar", Feb: "fewral", Mar: "mart", Apr: "aprel", May: "maý", Jun: "iýun", Jul: "iýul", Aug: "awgust", Sep: "sentýabr", Oct: "oktýabr", Nov: "noýabr", Dec: "dekabr"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "ýb", Mon: "db", Tue: "sb", Wed: "çb", Thu: "pb", Fri: "an", Sat: "şb"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "Ý", Mon: "D", Tue: "S", Wed: "Ç", Thu: "P", Fri: "A", Sat: "Ş"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "ýekşenbe", Mon: "duşenbe", Tue: "sişenbe", Wed: "çarşenbe", Thu: "penşenbe", Fri: "anna", Sat: "şenbe"},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
