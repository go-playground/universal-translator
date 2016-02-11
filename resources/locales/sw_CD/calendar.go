package sw_CD

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "", Medium: "", Short: "d/M/y"},
			Time:     ut.CalendarDateFormat{},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "mkw", Feb: "mpi", Mar: "mtu", Apr: "min", May: "mtn", Jun: "mst", Jul: "msb", Aug: "mun", Sep: "mts", Oct: "mku", Nov: "mkm", Dec: "mkb"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "k", Feb: "p", Mar: "t", Apr: "i", May: "t", Jun: "s", Jul: "s", Aug: "m", Sep: "t", Oct: "k", Nov: "m", Dec: "m"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "mwezi ya kwanja", Feb: "mwezi ya pili", Mar: "mwezi ya tatu", Apr: "mwezi ya ine", May: "mwezi ya tanu", Jun: "mwezi ya sita", Jul: "mwezi ya saba", Aug: "mwezi ya munane", Sep: "mwezi ya tisa", Oct: "mwezi ya kumi", Nov: "mwezi ya kumi na moya", Dec: "mwezi ya kumi ya mbili"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "yen", Mon: "kwa", Tue: "pil", Wed: "tat", Thu: "ine", Fri: "tan", Sat: "sit"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "y", Mon: "k", Tue: "p", Wed: "t", Thu: "i", Fri: "t", Sat: "s"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "siku ya yenga", Mon: "siku ya kwanza", Tue: "siku ya pili", Wed: "siku ya tatu", Thu: "siku ya ine", Fri: "siku ya tanu", Sat: "siku ya sita"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "ya asubuyi", PM: "ya muchana"},
			},
		},
	}
}
