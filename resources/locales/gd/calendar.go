package gd

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d'mh' MMMM y", Long: "d'mh' MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Faoi", Feb: "Gearr", Mar: "Màrt", Apr: "Gibl", May: "Cèit", Jun: "Ògmh", Jul: "Iuch", Aug: "Lùna", Sep: "Sult", Oct: "Dàmh", Nov: "Samh", Dec: "Dùbh"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "F", Feb: "G", Mar: "M", Apr: "G", May: "C", Jun: "Ò", Jul: "I", Aug: "L", Sep: "S", Oct: "D", Nov: "S", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Am Faoilleach", Feb: "An Gearran", Mar: "Am Màrt", Apr: "An Giblean", May: "An Cèitean", Jun: "An t-Ògmhios", Jul: "An t-Iuchar", Aug: "An Lùnastal", Sep: "An t-Sultain", Oct: "An Dàmhair", Nov: "An t-Samhain", Dec: "An Dùbhlachd"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "DiD", Mon: "DiL", Tue: "DiM", Wed: "DiC", Thu: "Dia", Fri: "Dih", Sat: "DiS"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "C", Thu: "A", Fri: "H", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "Dò", Mon: "Lu", Tue: "Mà", Wed: "Ci", Thu: "Da", Fri: "hA", Sat: "Sa"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "DiDòmhnaich", Mon: "DiLuain", Tue: "DiMàirt", Wed: "DiCiadain", Thu: "DiarDaoin", Fri: "DihAoine", Sat: "DiSathairne"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "m", PM: "f"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "m", PM: "f"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "m", PM: "f"},
			},
		},
	}
}
