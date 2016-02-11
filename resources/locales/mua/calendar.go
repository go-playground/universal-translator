package mua

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "FLO", Feb: "CLA", Mar: "CKI", Apr: "FMF", May: "MAD", Jun: "MBI", Jul: "MLI", Aug: "MAM", Sep: "FDE", Oct: "FMU", Nov: "FGW", Dec: "FYU"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "O", Feb: "A", Mar: "I", Apr: "F", May: "D", Jun: "B", Jul: "L", Aug: "M", Sep: "E", Oct: "U", Nov: "W", Dec: "Y"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Fĩi Loo", Feb: "Cokcwaklaŋne", Mar: "Cokcwaklii", Apr: "Fĩi Marfoo", May: "Madǝǝuutǝbijaŋ", Jun: "Mamǝŋgwãafahbii", Jul: "Mamǝŋgwãalii", Aug: "Madǝmbii", Sep: "Fĩi Dǝɓlii", Oct: "Fĩi Mundaŋ", Nov: "Fĩi Gwahlle", Dec: "Fĩi Yuru"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Cya", Mon: "Cla", Tue: "Czi", Wed: "Cko", Thu: "Cka", Fri: "Cga", Sat: "Cze"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "Y", Mon: "L", Tue: "Z", Wed: "O", Thu: "A", Fri: "G", Sat: "E"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Com’yakke", Mon: "Comlaaɗii", Tue: "Comzyiiɗii", Wed: "Comkolle", Thu: "Comkaldǝɓlii", Fri: "Comgaisuu", Sat: "Comzyeɓsuu"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "comme", PM: "lilli"},
			},
		},
	}
}
