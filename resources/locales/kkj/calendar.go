package kkj

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE dd MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM y"},
			Time:     ut.CalendarDateFormat{Full: "", Long: "", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "pamba", Feb: "wanja", Mar: "mbiyɔ mɛndoŋgɔ", Apr: "Nyɔlɔmbɔŋgɔ", May: "Mɔnɔ ŋgbanja", Jun: "Nyaŋgwɛ ŋgbanja", Jul: "kuŋgwɛ", Aug: "fɛ", Sep: "njapi", Oct: "nyukul", Nov: "11", Dec: "ɓulɓusɛ"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "sɔndi", Mon: "lundi", Tue: "mardi", Wed: "mɛrkɛrɛdi", Thu: "yedi", Fri: "vaŋdɛrɛdi", Sat: "mɔnɔ sɔndi"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "so", Mon: "lu", Tue: "ma", Wed: "mɛ", Thu: "ye", Fri: "va", Sat: "ms"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "so", Mon: "lu", Tue: "ma", Wed: "mɛ", Thu: "ye", Fri: "va", Sat: "ms"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "sɔndi", Mon: "lundi", Tue: "mardi", Wed: "mɛrkɛrɛdi", Thu: "yedi", Fri: "vaŋdɛrɛdi", Sat: "mɔnɔ sɔndi"},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
