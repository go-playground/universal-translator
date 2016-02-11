package et

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM y", Short: "dd.MM.yy"},
			Time:     ut.CalendarDateFormat{Full: "H:mm.ss zzzz", Long: "H:mm.ss z", Medium: "H:mm.ss", Short: "H:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "jaan", Feb: "veebr", Mar: "märts", Apr: "apr", May: "mai", Jun: "juuni", Jul: "juuli", Aug: "aug", Sep: "sept", Oct: "okt", Nov: "nov", Dec: "dets"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "V", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "jaanuar", Feb: "veebruar", Mar: "märts", Apr: "aprill", May: "mai", Jun: "juuni", Jul: "juuli", Aug: "august", Sep: "september", Oct: "oktoober", Nov: "november", Dec: "detsember"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "P", Mon: "E", Tue: "T", Wed: "K", Thu: "N", Fri: "R", Sat: "L"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "P", Mon: "E", Tue: "T", Wed: "K", Thu: "N", Fri: "R", Sat: "L"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "P", Mon: "E", Tue: "T", Wed: "K", Thu: "N", Fri: "R", Sat: "L"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "pühapäev", Mon: "esmaspäev", Tue: "teisipäev", Wed: "kolmapäev", Thu: "neljapäev", Fri: "reede", Sat: "laupäev"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			},
		},
	}
}
