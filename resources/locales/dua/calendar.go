package dua

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "di", Feb: "ŋgɔn", Mar: "sɔŋ", Apr: "diɓ", May: "emi", Jun: "esɔ", Jul: "mad", Aug: "diŋ", Sep: "nyɛt", Oct: "may", Nov: "tin", Dec: "elá"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "d", Feb: "ŋ", Mar: "s", Apr: "d", May: "e", Jun: "e", Jul: "m", Aug: "d", Sep: "n", Oct: "m", Nov: "t", Dec: "e"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "dimɔ́di", Feb: "ŋgɔndɛ", Mar: "sɔŋɛ", Apr: "diɓáɓá", May: "emiasele", Jun: "esɔpɛsɔpɛ", Jul: "madiɓɛ́díɓɛ́", Aug: "diŋgindi", Sep: "nyɛtɛki", Oct: "mayésɛ́", Nov: "tiníní", Dec: "eláŋgɛ́"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "ét", Mon: "mɔ́s", Tue: "kwa", Wed: "muk", Thu: "ŋgi", Fri: "ɗón", Sat: "esa"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "e", Mon: "m", Tue: "k", Wed: "m", Thu: "ŋ", Fri: "ɗ", Sat: "e"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "éti", Mon: "mɔ́sú", Tue: "kwasú", Wed: "mukɔ́sú", Thu: "ŋgisú", Fri: "ɗónɛsú", Sat: "esaɓasú"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "idiɓa", PM: "ebyámu"},
			},
		},
	}
}
