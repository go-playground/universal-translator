package az

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "d MMMM y, EEEE", Long: "d MMMM y", Medium: "d MMM y", Short: "dd.MM.yy"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "yan", Feb: "fev", Mar: "mar", Apr: "apr", May: "may", Jun: "iyn", Jul: "iyl", Aug: "avq", Sep: "sen", Oct: "okt", Nov: "noy", Dec: "dek"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Yanvar", Feb: "Fevral", Mar: "Mart", Apr: "Aprel", May: "May", Jun: "İyun", Jul: "İyul", Aug: "Avqust", Sep: "Sentyabr", Oct: "Oktyabr", Nov: "Noyabr", Dec: "Dekabr"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "B.", Mon: "B.E.", Tue: "Ç.A.", Wed: "Ç.", Thu: "C.A.", Fri: "C.", Sat: "Ş."},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "7", Mon: "1", Tue: "2", Wed: "3", Thu: "4", Fri: "5", Sat: "6"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "B.", Mon: "B.E.", Tue: "Ç.A.", Wed: "Ç.", Thu: "C.A.", Fri: "C.", Sat: "Ş."},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "bazar", Mon: "bazar ertəsi", Tue: "çərşənbə axşamı", Wed: "çərşənbə", Thu: "cümə axşamı", Fri: "cümə", Sat: "şənbə"},
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
