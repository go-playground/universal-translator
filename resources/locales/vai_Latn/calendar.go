package vai_Latn

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "luukao kemã", Feb: "ɓandaɓu", Mar: "vɔɔ", Apr: "fulu", May: "goo", Jun: "6", Jul: "7", Aug: "kɔnde", Sep: "saah", Oct: "galo", Nov: "kenpkato ɓololɔ", Dec: "luukao lɔma"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{},
				Narrow:      ut.CalendarDayFormatNameValue{},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "lahadi", Mon: "tɛɛnɛɛ", Tue: "talata", Wed: "alaba", Thu: "aimisa", Fri: "aijima", Sat: "siɓiti"},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
