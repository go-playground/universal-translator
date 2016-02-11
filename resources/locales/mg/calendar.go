package mg

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mar", Apr: "Apr", May: "Mey", Jun: "Jon", Jul: "Jol", Aug: "Aog", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Des"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Janoary", Feb: "Febroary", Mar: "Martsa", Apr: "Aprily", May: "Mey", Jun: "Jona", Jul: "Jolay", Aug: "Aogositra", Sep: "Septambra", Oct: "Oktobra", Nov: "Novambra", Dec: "Desambra"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Alah", Mon: "Alats", Tue: "Tal", Wed: "Alar", Thu: "Alak", Fri: "Zom", Sat: "Asab"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "A", Mon: "A", Tue: "T", Wed: "A", Thu: "A", Fri: "Z", Sat: "A"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Alahady", Mon: "Alatsinainy", Tue: "Talata", Wed: "Alarobia", Thu: "Alakamisy", Fri: "Zoma", Sat: "Asabotsy"},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
