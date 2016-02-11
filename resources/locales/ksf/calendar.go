package ksf

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "ŋ1", Feb: "ŋ2", Mar: "ŋ3", Apr: "ŋ4", May: "ŋ5", Jun: "ŋ6", Jul: "ŋ7", Aug: "ŋ8", Sep: "ŋ9", Oct: "ŋ10", Nov: "ŋ11", Dec: "ŋ12"},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "ŋwíí a ntɔ́ntɔ", Feb: "ŋwíí akǝ bɛ́ɛ", Mar: "ŋwíí akǝ ráá", Apr: "ŋwíí akǝ nin", May: "ŋwíí akǝ táan", Jun: "ŋwíí akǝ táafɔk", Jul: "ŋwíí akǝ táabɛɛ", Aug: "ŋwíí akǝ táaraa", Sep: "ŋwíí akǝ táanin", Oct: "ŋwíí akǝ ntɛk", Nov: "ŋwíí akǝ ntɛk di bɔ́k", Dec: "ŋwíí akǝ ntɛk di bɛ́ɛ"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "sɔ́n", Mon: "lǝn", Tue: "maa", Wed: "mɛk", Thu: "jǝǝ", Fri: "júm", Sat: "sam"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "s", Mon: "l", Tue: "m", Wed: "m", Thu: "j", Fri: "j", Sat: "s"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "sɔ́ndǝ", Mon: "lǝndí", Tue: "maadí", Wed: "mɛkrɛdí", Thu: "jǝǝdí", Fri: "júmbá", Sat: "samdí"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "sárúwá", PM: "cɛɛ́nko"},
			},
		},
	}
}
