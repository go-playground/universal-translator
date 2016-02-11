package ig

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Jen", Feb: "Feb", Mar: "Maa", Apr: "Epr", May: "Mee", Jun: "Juu", Jul: "Jul", Aug: "Ọgọ", Sep: "Sep", Oct: "Ọkt", Nov: "Nov", Dec: "Dis"},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Jenụwarị", Feb: "Febrụwarị", Mar: "Maachị", Apr: "Eprel", May: "Mee", Jun: "Juun", Jul: "Julaị", Aug: "Ọgọọst", Sep: "Septemba", Oct: "Ọktoba", Nov: "Novemba", Dec: "Disemba"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Ụka", Mon: "Mọn", Tue: "Tiu", Wed: "Wen", Thu: "Tọọ", Fri: "Fraị", Sat: "Sat"},
				Narrow:      ut.CalendarDayFormatNameValue{},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Mbọsị Ụka", Mon: "Mọnde", Tue: "Tiuzdee", Wed: "Wenezdee", Thu: "Tọọzdee", Fri: "Fraịdee", Sat: "Satọdee"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "A.M.", PM: "P.M."},
			},
		},
	}
}
