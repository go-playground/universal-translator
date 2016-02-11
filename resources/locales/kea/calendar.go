package kea

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d 'di' MMMM 'di' y", Long: "d 'di' MMMM 'di' y", Medium: "d MMM y", Short: "d/M/y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mar", Apr: "Abr", May: "Mai", Jun: "Jun", Jul: "Jul", Aug: "Ago", Sep: "Set", Oct: "Otu", Nov: "Nuv", Dec: "Diz"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Janeru", Feb: "Febreru", Mar: "Marsu", Apr: "Abril", May: "Maiu", Jun: "Junhu", Jul: "Julhu", Aug: "Agostu", Sep: "Setenbru", Oct: "Otubru", Nov: "Nuvenbru", Dec: "Dizenbru"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "dum", Mon: "sig", Tue: "ter", Wed: "kua", Thu: "kin", Fri: "ses", Sat: "sab"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "d", Mon: "s", Tue: "t", Wed: "k", Thu: "k", Fri: "s", Sat: "s"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "du", Mon: "si", Tue: "te", Wed: "ku", Thu: "ki", Fri: "se", Sat: "sa"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "dumingu", Mon: "sigunda-fera", Tue: "tersa-fera", Wed: "kuarta-fera", Thu: "kinta-fera", Fri: "sesta-fera", Sat: "sabadu"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "am", PM: "pm"},
			},
		},
	}
}
