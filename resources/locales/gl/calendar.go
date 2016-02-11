package gl

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE dd MMMM y", Long: "dd MMMM y", Medium: "d MMM, y", Short: "dd/MM/yy"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Xan", Feb: "Feb", Mar: "Mar", Apr: "Abr", May: "Mai", Jun: "Xuñ", Jul: "Xul", Aug: "Ago", Sep: "Set", Oct: "Out", Nov: "Nov", Dec: "Dec"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "X", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "X", Jul: "X", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Xaneiro", Feb: "Febreiro", Mar: "Marzo", Apr: "Abril", May: "Maio", Jun: "Xuño", Jul: "Xullo", Aug: "Agosto", Sep: "Setembro", Oct: "Outubro", Nov: "Novembro", Dec: "Decembro"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Dom", Mon: "Lun", Tue: "Mar", Wed: "Mér", Thu: "Xov", Fri: "Ven", Sat: "Sáb"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "M", Thu: "X", Fri: "V", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "Dom", Mon: "Luns", Tue: "Mt", Wed: "Mc", Thu: "Xv", Fri: "Ven", Sat: "Sáb"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Domingo", Mon: "Luns", Tue: "Martes", Wed: "Mércores", Thu: "Xoves", Fri: "Venres", Sat: "Sábado"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
			},
		},
	}
}
