package luo

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "DAC", Feb: "DAR", Mar: "DAD", Apr: "DAN", May: "DAH", Jun: "DAU", Jul: "DAO", Aug: "DAB", Sep: "DOC", Oct: "DAP", Nov: "DGI", Dec: "DAG"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "C", Feb: "R", Mar: "D", Apr: "N", May: "B", Jun: "U", Jul: "B", Aug: "B", Sep: "C", Oct: "P", Nov: "C", Dec: "P"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Dwe mar Achiel", Feb: "Dwe mar Ariyo", Mar: "Dwe mar Adek", Apr: "Dwe mar Ang’wen", May: "Dwe mar Abich", Jun: "Dwe mar Auchiel", Jul: "Dwe mar Abiriyo", Aug: "Dwe mar Aboro", Sep: "Dwe mar Ochiko", Oct: "Dwe mar Apar", Nov: "Dwe mar gi achiel", Dec: "Dwe mar Apar gi ariyo"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "JMP", Mon: "WUT", Tue: "TAR", Wed: "TAD", Thu: "TAN", Fri: "TAB", Sat: "NGS"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "J", Mon: "W", Tue: "T", Wed: "T", Thu: "T", Fri: "T", Sat: "N"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Jumapil", Mon: "Wuok Tich", Tue: "Tich Ariyo", Wed: "Tich Adek", Thu: "Tich Ang’wen", Fri: "Tich Abich", Sat: "Ngeso"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "OD", PM: "OT"},
			},
		},
	}
}
