package mgh

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Kwa", Feb: "Una", Mar: "Rar", Apr: "Che", May: "Tha", Jun: "Moc", Jul: "Sab", Aug: "Nan", Sep: "Tis", Oct: "Kum", Nov: "Moj", Dec: "Yel"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "K", Feb: "U", Mar: "R", Apr: "C", May: "T", Jun: "M", Jul: "S", Aug: "N", Sep: "T", Oct: "K", Nov: "M", Dec: "Y"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Mweri wo kwanza", Feb: "Mweri wo unayeli", Mar: "Mweri wo uneraru", Apr: "Mweri wo unecheshe", May: "Mweri wo unethanu", Jun: "Mweri wo thanu na mocha", Jul: "Mweri wo saba", Aug: "Mweri wo nane", Sep: "Mweri wo tisa", Oct: "Mweri wo kumi", Nov: "Mweri wo kumi na moja", Dec: "Mweri wo kumi na yel’li"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Sab", Mon: "Jtt", Tue: "Jnn", Wed: "Jtn", Thu: "Ara", Fri: "Iju", Sat: "Jmo"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "S", Mon: "J", Tue: "J", Wed: "J", Thu: "A", Fri: "I", Sat: "J"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Sabato", Mon: "Jumatatu", Tue: "Jumanne", Wed: "Jumatano", Thu: "Arahamisi", Fri: "Ijumaa", Sat: "Jumamosi"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "wichishu", PM: "mchochil’l"},
			},
		},
	}
}
