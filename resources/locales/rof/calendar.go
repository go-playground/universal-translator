package rof

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "M1", Feb: "M2", Mar: "M3", Apr: "M4", May: "M5", Jun: "M6", Jul: "M7", Aug: "M8", Sep: "M9", Oct: "M10", Nov: "M11", Dec: "M12"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "K", Feb: "K", Mar: "K", Apr: "K", May: "T", Jun: "S", Jul: "S", Aug: "N", Sep: "T", Oct: "I", Nov: "I", Dec: "I"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Mweri wa kwanza", Feb: "Mweri wa kaili", Mar: "Mweri wa katatu", Apr: "Mweri wa kaana", May: "Mweri wa tanu", Jun: "Mweri wa sita", Jul: "Mweri wa saba", Aug: "Mweri wa nane", Sep: "Mweri wa tisa", Oct: "Mweri wa ikumi", Nov: "Mweri wa ikumi na moja", Dec: "Mweri wa ikumi na mbili"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Ijp", Mon: "Ijt", Tue: "Ijn", Wed: "Ijtn", Thu: "Alh", Fri: "Iju", Sat: "Ijm"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "2", Mon: "3", Tue: "4", Wed: "5", Thu: "6", Fri: "7", Sat: "1"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Ijumapili", Mon: "Ijumatatu", Tue: "Ijumanne", Wed: "Ijumatano", Thu: "Alhamisi", Fri: "Ijumaa", Sat: "Ijumamosi"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "kang’ama", PM: "kingoto"},
			},
		},
	}
}
