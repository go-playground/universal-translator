package kln

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Mul", Feb: "Ngat", Mar: "Taa", Apr: "Iwo", May: "Mam", Jun: "Paa", Jul: "Nge", Aug: "Roo", Sep: "Bur", Oct: "Epe", Nov: "Kpt", Dec: "Kpa"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "M", Feb: "N", Mar: "T", Apr: "I", May: "M", Jun: "P", Jul: "N", Aug: "R", Sep: "B", Oct: "E", Nov: "K", Dec: "K"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Mulgul", Feb: "Ng’atyaato", Mar: "Kiptaamo", Apr: "Iwootkuut", May: "Mamuut", Jun: "Paagi", Jul: "Ng’eiyeet", Aug: "Rooptui", Sep: "Bureet", Oct: "Epeeso", Nov: "Kipsuunde ne taai", Dec: "Kipsuunde nebo aeng’"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Kts", Mon: "Kot", Tue: "Koo", Wed: "Kos", Thu: "Koa", Fri: "Kom", Sat: "Kol"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "T", Mon: "T", Tue: "O", Wed: "S", Thu: "A", Fri: "M", Sat: "L"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Kotisap", Mon: "Kotaai", Tue: "Koaeng’", Wed: "Kosomok", Thu: "Koang’wan", Fri: "Komuut", Sat: "Kolo"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "krn", PM: "koosk"},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "karoon", PM: "kooskoliny"},
			},
		},
	}
}
