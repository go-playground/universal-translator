package mer

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "JAN", Feb: "FEB", Mar: "MAC", Apr: "ĨPU", May: "MĨĨ", Jun: "NJU", Jul: "NJR", Aug: "AGA", Sep: "SPT", Oct: "OKT", Nov: "NOV", Dec: "DEC"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "Ĩ", May: "M", Jun: "N", Jul: "N", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Januarĩ", Feb: "Feburuarĩ", Mar: "Machi", Apr: "Ĩpurũ", May: "Mĩĩ", Jun: "Njuni", Jul: "Njuraĩ", Aug: "Agasti", Sep: "Septemba", Oct: "Oktũba", Nov: "Novemba", Dec: "Dicemba"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "KIU", Mon: "MRA", Tue: "WAI", Wed: "WET", Thu: "WEN", Fri: "WTN", Sat: "JUM"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "K", Mon: "M", Tue: "W", Wed: "W", Thu: "W", Fri: "W", Sat: "J"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Kiumia", Mon: "Muramuko", Tue: "Wairi", Wed: "Wethatu", Thu: "Wena", Fri: "Wetano", Sat: "Jumamosi"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "RŨ", PM: "ŨG"},
			},
		},
	}
}
