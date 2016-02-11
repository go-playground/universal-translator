package ewo

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "ngo", Feb: "ngb", Mar: "ngl", Apr: "ngn", May: "ngt", Jun: "ngs", Jul: "ngz", Aug: "ngm", Sep: "nge", Oct: "nga", Nov: "ngad", Dec: "ngab"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "o", Feb: "b", Mar: "l", Apr: "n", May: "t", Jun: "s", Jul: "z", Aug: "m", Sep: "e", Oct: "a", Nov: "d", Dec: "b"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "ngɔn osú", Feb: "ngɔn bɛ̌", Mar: "ngɔn lála", Apr: "ngɔn nyina", May: "ngɔn tána", Jun: "ngɔn saməna", Jul: "ngɔn zamgbála", Aug: "ngɔn mwom", Sep: "ngɔn ebulú", Oct: "ngɔn awóm", Nov: "ngɔn awóm ai dziá", Dec: "ngɔn awóm ai bɛ̌"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "sɔ́n", Mon: "mɔ́n", Tue: "smb", Wed: "sml", Thu: "smn", Fri: "fúl", Sat: "sér"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "s", Mon: "m", Tue: "s", Wed: "s", Thu: "s", Fri: "f", Sat: "s"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "sɔ́ndɔ", Mon: "mɔ́ndi", Tue: "sɔ́ndɔ məlú mə́bɛ̌", Wed: "sɔ́ndɔ məlú mə́lɛ́", Thu: "sɔ́ndɔ məlú mə́nyi", Fri: "fúladé", Sat: "séradé"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "kíkíríg", PM: "ngəgógəle"},
			},
		},
	}
}
