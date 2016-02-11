package cs

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE d. MMMM y", Long: "d. MMMM y", Medium: "d. M. y", Short: "dd.MM.yy"},
			Time:     ut.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "led", Feb: "úno", Mar: "bře", Apr: "dub", May: "kvě", Jun: "čvn", Jul: "čvc", Aug: "srp", Sep: "zář", Oct: "říj", Nov: "lis", Dec: "pro"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "leden", Feb: "únor", Mar: "březen", Apr: "duben", May: "květen", Jun: "červen", Jul: "červenec", Aug: "srpen", Sep: "září", Oct: "říjen", Nov: "listopad", Dec: "prosinec"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "ne", Mon: "po", Tue: "út", Wed: "st", Thu: "čt", Fri: "pá", Sat: "so"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "N", Mon: "P", Tue: "Ú", Wed: "S", Thu: "Č", Fri: "P", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "ne", Mon: "po", Tue: "út", Wed: "st", Thu: "čt", Fri: "pá", Sat: "so"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "neděle", Mon: "pondělí", Tue: "úterý", Wed: "středa", Thu: "čtvrtek", Fri: "pátek", Sat: "sobota"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "dop.", PM: "odp."},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "dop.", PM: "odp."},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "dop.", PM: "odp."},
			},
		},
	}
}
