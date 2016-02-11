package dyo

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Sa", Feb: "Fe", Mar: "Ma", Apr: "Ab", May: "Me", Jun: "Su", Jul: "Sú", Aug: "Ut", Sep: "Se", Oct: "Ok", Nov: "No", Dec: "De"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "S", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "S", Jul: "S", Aug: "U", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Sanvie", Feb: "Fébirie", Mar: "Mars", Apr: "Aburil", May: "Mee", Jun: "Sueŋ", Jul: "Súuyee", Aug: "Ut", Sep: "Settembar", Oct: "Oktobar", Nov: "Novembar", Dec: "Disambar"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Dim", Mon: "Ten", Tue: "Tal", Wed: "Ala", Thu: "Ara", Fri: "Arj", Sat: "Sib"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "D", Mon: "T", Tue: "T", Wed: "A", Thu: "A", Fri: "A", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Dimas", Mon: "Teneŋ", Tue: "Talata", Wed: "Alarbay", Thu: "Aramisay", Fri: "Arjuma", Sat: "Sibiti"},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
