package shi

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"},
			Time:     ut.CalendarDateFormat{},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "ⵉⵏⵏ", Feb: "ⴱⵕⴰ", Mar: "ⵎⴰⵕ", Apr: "ⵉⴱⵔ", May: "ⵎⴰⵢ", Jun: "ⵢⵓⵏ", Jul: "ⵢⵓⵍ", Aug: "ⵖⵓⵛ", Sep: "ⵛⵓⵜ", Oct: "ⴽⵜⵓ", Nov: "ⵏⵓⵡ", Dec: "ⴷⵓⵊ"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "ⵉ", Feb: "ⴱ", Mar: "ⵎ", Apr: "ⵉ", May: "ⵎ", Jun: "ⵢ", Jul: "ⵢ", Aug: "ⵖ", Sep: "ⵛ", Oct: "ⴽ", Nov: "ⵏ", Dec: "ⴷ"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "ⵉⵏⵏⴰⵢⵔ", Feb: "ⴱⵕⴰⵢⵕ", Mar: "ⵎⴰⵕⵚ", Apr: "ⵉⴱⵔⵉⵔ", May: "ⵎⴰⵢⵢⵓ", Jun: "ⵢⵓⵏⵢⵓ", Jul: "ⵢⵓⵍⵢⵓⵣ", Aug: "ⵖⵓⵛⵜ", Sep: "ⵛⵓⵜⴰⵏⴱⵉⵔ", Oct: "ⴽⵜⵓⴱⵔ", Nov: "ⵏⵓⵡⴰⵏⴱⵉⵔ", Dec: "ⴷⵓⵊⴰⵏⴱⵉⵔ"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "ⴰⵙⴰ", Mon: "ⴰⵢⵏ", Tue: "ⴰⵙⵉ", Wed: "ⴰⴽⵕ", Thu: "ⴰⴽⵡ", Fri: "ⴰⵙⵉⵎ", Sat: "ⴰⵙⵉⴹ"},
				Narrow:      ut.CalendarDayFormatNameValue{},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "ⴰⵙⴰⵎⴰⵙ", Mon: "ⴰⵢⵏⴰⵙ", Tue: "ⴰⵙⵉⵏⴰⵙ", Wed: "ⴰⴽⵕⴰⵙ", Thu: "ⴰⴽⵡⴰⵙ", Fri: "ⵙⵉⵎⵡⴰⵙ", Sat: "ⴰⵙⵉⴹⵢⴰⵙ"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "ⵜⵉⴼⴰⵡⵜ", PM: "ⵜⴰⴷⴳⴳⵯⴰⵜ"},
			},
		},
	}
}
