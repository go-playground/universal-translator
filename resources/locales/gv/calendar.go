package gv

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE dd MMMM y", Long: "dd MMMM y", Medium: "MMM dd, y", Short: "dd/MM/yy"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "J-guer", Feb: "T-arree", Mar: "Mayrnt", Apr: "Avrril", May: "Boaldyn", Jun: "M-souree", Jul: "J-souree", Aug: "Luanistyn", Sep: "M-fouyir", Oct: "J-fouyir", Nov: "M-Houney", Dec: "M-Nollick"},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Jerrey-geuree", Feb: "Toshiaght-arree", Mar: "Mayrnt", Apr: "Averil", May: "Boaldyn", Jun: "Mean-souree", Jul: "Jerrey-souree", Aug: "Luanistyn", Sep: "Mean-fouyir", Oct: "Jerrey-fouyir", Nov: "Mee Houney", Dec: "Mee ny Nollick"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Jed", Mon: "Jel", Tue: "Jem", Wed: "Jerc", Thu: "Jerd", Fri: "Jeh", Sat: "Jes"},
				Narrow:      ut.CalendarDayFormatNameValue{},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Jedoonee", Mon: "Jelhein", Tue: "Jemayrt", Wed: "Jercean", Thu: "Jerdein", Fri: "Jeheiney", Sat: "Jesarn"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
			},
		},
	}
}
