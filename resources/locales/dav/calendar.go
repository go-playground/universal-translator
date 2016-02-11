package dav

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Imb", Feb: "Kaw", Mar: "Kad", Apr: "Kan", May: "Kas", Jun: "Kar", Jul: "Mfu", Aug: "Wun", Sep: "Ike", Oct: "Iku", Nov: "Imw", Dec: "Iwi"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "I", Feb: "K", Mar: "K", Apr: "K", May: "K", Jun: "K", Jul: "M", Aug: "W", Sep: "I", Oct: "I", Nov: "I", Dec: "I"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Mori ghwa imbiri", Feb: "Mori ghwa kawi", Mar: "Mori ghwa kadadu", Apr: "Mori ghwa kana", May: "Mori ghwa kasanu", Jun: "Mori ghwa karandadu", Jul: "Mori ghwa mfungade", Aug: "Mori ghwa wunyanya", Sep: "Mori ghwa ikenda", Oct: "Mori ghwa ikumi", Nov: "Mori ghwa ikumi na imweri", Dec: "Mori ghwa ikumi na iwi"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Jum", Mon: "Jim", Tue: "Kaw", Wed: "Kad", Thu: "Kan", Fri: "Kas", Sat: "Ngu"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "J", Mon: "J", Tue: "K", Wed: "K", Thu: "K", Fri: "K", Sat: "N"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Ituku ja jumwa", Mon: "Kuramuka jimweri", Tue: "Kuramuka kawi", Wed: "Kuramuka kadadu", Thu: "Kuramuka kana", Fri: "Kuramuka kasanu", Sat: "Kifula nguwo"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "Luma lwa K", PM: "luma lwa p"},
			},
		},
	}
}
