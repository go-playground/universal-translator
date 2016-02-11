package cy

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/yy"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} 'am' {0}", Long: "{1} 'am' {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Ion", Feb: "Chw", Mar: "Maw", Apr: "Ebr", May: "Mai", Jun: "Meh", Jul: "Gor", Aug: "Awst", Sep: "Medi", Oct: "Hyd", Nov: "Tach", Dec: "Rhag"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "I", Feb: "Ch", Mar: "M", Apr: "E", May: "M", Jun: "M", Jul: "G", Aug: "A", Sep: "M", Oct: "H", Nov: "T", Dec: "Rh"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Ionawr", Feb: "Chwefror", Mar: "Mawrth", Apr: "Ebrill", May: "Mai", Jun: "Mehefin", Jul: "Gorffennaf", Aug: "Awst", Sep: "Medi", Oct: "Hydref", Nov: "Tachwedd", Dec: "Rhagfyr"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Sul", Mon: "Llun", Tue: "Maw", Wed: "Mer", Thu: "Iau", Fri: "Gwe", Sat: "Sad"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "S", Mon: "Ll", Tue: "M", Wed: "M", Thu: "I", Fri: "G", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "Su", Mon: "Ll", Tue: "Ma", Wed: "Me", Thu: "Ia", Fri: "Gw", Sat: "Sa"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Dydd Sul", Mon: "Dydd Llun", Tue: "Dydd Mawrth", Wed: "Dydd Mercher", Thu: "Dydd Iau", Fri: "Dydd Gwener", Sat: "Dydd Sadwrn"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			},
		},
	}
}
