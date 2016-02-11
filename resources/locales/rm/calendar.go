package rm

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, 'ils' d 'da' MMMM y", Long: "d 'da' MMMM y", Medium: "dd-MM-y", Short: "dd-MM-yy"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "schan.", Feb: "favr.", Mar: "mars", Apr: "avr.", May: "matg", Jun: "zercl.", Jul: "fan.", Aug: "avust", Sep: "sett.", Oct: "oct.", Nov: "nov.", Dec: "dec."},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "S", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "Z", Jul: "F", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "schaner", Feb: "favrer", Mar: "mars", Apr: "avrigl", May: "matg", Jun: "zercladur", Jul: "fanadur", Aug: "avust", Sep: "settember", Oct: "october", Nov: "november", Dec: "december"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "du", Mon: "gli", Tue: "ma", Wed: "me", Thu: "gie", Fri: "ve", Sat: "so"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "D", Mon: "G", Tue: "M", Wed: "M", Thu: "G", Fri: "V", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "dumengia", Mon: "glindesdi", Tue: "mardi", Wed: "mesemna", Thu: "gievgia", Fri: "venderdi", Sat: "sonda"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "am", PM: "sm"},
			},
		},
	}
}
