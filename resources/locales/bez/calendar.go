package bez

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Hut", Feb: "Vil", Mar: "Dat", Apr: "Tai", May: "Han", Jun: "Sit", Jul: "Sab", Aug: "Nan", Sep: "Tis", Oct: "Kum", Nov: "Kmj", Dec: "Kmb"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "H", Feb: "V", Mar: "D", Apr: "T", May: "H", Jun: "S", Jul: "S", Aug: "N", Sep: "T", Oct: "K", Nov: "K", Dec: "K"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "pa mwedzi gwa hutala", Feb: "pa mwedzi gwa wuvili", Mar: "pa mwedzi gwa wudatu", Apr: "pa mwedzi gwa wutai", May: "pa mwedzi gwa wuhanu", Jun: "pa mwedzi gwa sita", Jul: "pa mwedzi gwa saba", Aug: "pa mwedzi gwa nane", Sep: "pa mwedzi gwa tisa", Oct: "pa mwedzi gwa kumi", Nov: "pa mwedzi gwa kumi na moja", Dec: "pa mwedzi gwa kumi na mbili"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Mul", Mon: "Vil", Tue: "Hiv", Wed: "Hid", Thu: "Hit", Fri: "Hih", Sat: "Lem"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "M", Mon: "J", Tue: "H", Wed: "H", Thu: "H", Fri: "W", Sat: "J"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "pa mulungu", Mon: "pa shahuviluha", Tue: "pa hivili", Wed: "pa hidatu", Thu: "pa hitayi", Fri: "pa hihanu", Sat: "pa shahulembela"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "pamilau", PM: "pamunyi"},
			},
		},
	}
}
