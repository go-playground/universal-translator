package wae

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM y", Short: ""},
			Time:     ut.CalendarDateFormat{},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Jen", Feb: "Hor", Mar: "Mär", Apr: "Abr", May: "Mei", Jun: "Brá", Jul: "Hei", Aug: "Öig", Sep: "Her", Oct: "Wím", Nov: "Win", Dec: "Chr"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "H", Mar: "M", Apr: "A", May: "M", Jun: "B", Jul: "H", Aug: "Ö", Sep: "H", Oct: "W", Nov: "W", Dec: "C"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Jenner", Feb: "Hornig", Mar: "Märze", Apr: "Abrille", May: "Meije", Jun: "Bráčet", Jul: "Heiwet", Aug: "Öigšte", Sep: "Herbštmánet", Oct: "Wímánet", Nov: "Wintermánet", Dec: "Chrištmánet"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Sun", Mon: "Män", Tue: "Ziš", Wed: "Mit", Thu: "Fró", Fri: "Fri", Sat: "Sam"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "Z", Wed: "M", Thu: "F", Fri: "F", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Sunntag", Mon: "Mäntag", Tue: "Zištag", Wed: "Mittwuč", Thu: "Fróntag", Fri: "Fritag", Sat: "Samštag"},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
