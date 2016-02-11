package shi_Latn

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "inn", Feb: "bṛa", Mar: "maṛ", Apr: "ibr", May: "may", Jun: "yun", Jul: "yul", Aug: "ɣuc", Sep: "cut", Oct: "ktu", Nov: "nuw", Dec: "duj"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "i", Feb: "b", Mar: "m", Apr: "i", May: "m", Jun: "y", Jul: "y", Aug: "ɣ", Sep: "c", Oct: "k", Nov: "n", Dec: "d"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "innayr", Feb: "bṛayṛ", Mar: "maṛṣ", Apr: "ibrir", May: "mayyu", Jun: "yunyu", Jul: "yulyuz", Aug: "ɣuct", Sep: "cutanbir", Oct: "ktubr", Nov: "nuwanbir", Dec: "dujanbir"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "asa", Mon: "ayn", Tue: "asi", Wed: "akṛ", Thu: "akw", Fri: "asim", Sat: "asiḍ"},
				Narrow:      ut.CalendarDayFormatNameValue{},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "asamas", Mon: "aynas", Tue: "asinas", Wed: "akṛas", Thu: "akwas", Fri: "asimwas", Sat: "asiḍyas"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "tifawt", PM: "tadggʷat"},
			},
		},
	}
}
