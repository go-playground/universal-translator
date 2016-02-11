package hsb

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d.M.y", Short: "d.M.yy"},
			Time:     ut.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm 'hodź'."},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "jan", Feb: "feb", Mar: "měr", Apr: "apr", May: "mej", Jun: "jun", Jul: "jul", Aug: "awg", Sep: "sep", Oct: "okt", Nov: "now", Dec: "dec"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "j", Feb: "f", Mar: "m", Apr: "a", May: "m", Jun: "j", Jul: "j", Aug: "a", Sep: "s", Oct: "o", Nov: "n", Dec: "d"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "januar", Feb: "februar", Mar: "měrc", Apr: "apryl", May: "meja", Jun: "junij", Jul: "julij", Aug: "awgust", Sep: "september", Oct: "oktober", Nov: "nowember", Dec: "december"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "nje", Mon: "pón", Tue: "wut", Wed: "srj", Thu: "štw", Fri: "pja", Sat: "sob"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "n", Mon: "p", Tue: "w", Wed: "s", Thu: "š", Fri: "p", Sat: "s"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "nj", Mon: "pó", Tue: "wu", Wed: "sr", Thu: "št", Fri: "pj", Sat: "so"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "njedźela", Mon: "póndźela", Tue: "wutora", Wed: "srjeda", Thu: "štwórtk", Fri: "pjatk", Sat: "sobota"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "dop.", PM: "pop."},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "dopołdnja", PM: "popołdnju"},
			},
		},
	}
}
