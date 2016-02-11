package mt

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d 'ta'’ MMMM y", Long: "d 'ta'’ MMMM y", Medium: "dd MMM y", Short: "dd/MM/y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Fra", Mar: "Mar", Apr: "Apr", May: "Mej", Jun: "Ġun", Jul: "Lul", Aug: "Aww", Sep: "Set", Oct: "Ott", Nov: "Nov", Dec: "Diċ"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "Jn", Feb: "Fr", Mar: "Mz", Apr: "Ap", May: "Mj", Jun: "Ġn", Jul: "Lj", Aug: "Aw", Sep: "St", Oct: "Ob", Nov: "Nv", Dec: "Dċ"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Jannar", Feb: "Frar", Mar: "Marzu", Apr: "April", May: "Mejju", Jun: "Ġunju", Jul: "Lulju", Aug: "Awwissu", Sep: "Settembru", Oct: "Ottubru", Nov: "Novembru", Dec: "Diċembru"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Ħad", Mon: "Tne", Tue: "Tli", Wed: "Erb", Thu: "Ħam", Fri: "Ġim", Sat: "Sib"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "Ħd", Mon: "Tn", Tue: "Tl", Wed: "Er", Thu: "Ħm", Fri: "Ġm", Sat: "Sb"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Il-Ħadd", Mon: "It-Tnejn", Tue: "It-Tlieta", Wed: "L-Erbgħa", Thu: "Il-Ħamis", Fri: "Il-Ġimgħa", Sat: "Is-Sibt"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			},
		},
	}
}
