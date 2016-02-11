package sg

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Nye", Feb: "Ful", Mar: "Mbä", Apr: "Ngu", May: "Bêl", Jun: "Fön", Jul: "Len", Aug: "Kük", Sep: "Mvu", Oct: "Ngb", Nov: "Nab", Dec: "Kak"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "N", Feb: "F", Mar: "M", Apr: "N", May: "B", Jun: "F", Jul: "L", Aug: "K", Sep: "M", Oct: "N", Nov: "N", Dec: "K"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Nyenye", Feb: "Fulundïgi", Mar: "Mbängü", Apr: "Ngubùe", May: "Bêläwü", Jun: "Föndo", Jul: "Lengua", Aug: "Kükürü", Sep: "Mvuka", Oct: "Ngberere", Nov: "Nabändüru", Dec: "Kakauka"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Bk1", Mon: "Bk2", Tue: "Bk3", Wed: "Bk4", Thu: "Bk5", Fri: "Lâp", Sat: "Lây"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "K", Mon: "S", Tue: "T", Wed: "S", Thu: "K", Fri: "P", Sat: "Y"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Bikua-ôko", Mon: "Bïkua-ûse", Tue: "Bïkua-ptâ", Wed: "Bïkua-usïö", Thu: "Bïkua-okü", Fri: "Lâpôsö", Sat: "Lâyenga"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "ND", PM: "LK"},
			},
		},
	}
}
