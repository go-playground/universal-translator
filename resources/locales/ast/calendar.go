package ast

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM 'de' y", Long: "d MMMM 'de' y", Medium: "d MMM y", Short: "d/M/yy"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} 'a' 'les' {0}", Long: "{1} 'a' 'les' {0}", Medium: "{1}, {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Xin", Feb: "Feb", Mar: "Mar", Apr: "Abr", May: "May", Jun: "Xun", Jul: "Xnt", Aug: "Ago", Sep: "Set", Oct: "Och", Nov: "Pay", Dec: "Avi"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "X", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "X", Jul: "X", Aug: "A", Sep: "S", Oct: "O", Nov: "P", Dec: "A"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "xineru", Feb: "febreru", Mar: "marzu", Apr: "abril", May: "mayu", Jun: "xunu", Jul: "xunetu", Aug: "agostu", Sep: "setiembre", Oct: "ochobre", Nov: "payares", Dec: "avientu"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "dom", Mon: "llu", Tue: "mar", Wed: "mié", Thu: "xue", Fri: "vie", Sat: "sáb"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "M", Thu: "X", Fri: "V", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "do", Mon: "ll", Tue: "ma", Wed: "mi", Thu: "xu", Fri: "vi", Sat: "sá"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "domingu", Mon: "llunes", Tue: "martes", Wed: "miércoles", Thu: "xueves", Fri: "vienres", Sat: "sábadu"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "mañana", PM: "tardi"},
			},
		},
	}
}
