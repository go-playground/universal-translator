package ln

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "yan", Feb: "fbl", Mar: "msi", Apr: "apl", May: "mai", Jun: "yun", Jul: "yul", Aug: "agt", Sep: "stb", Oct: "ɔtb", Nov: "nvb", Dec: "dsb"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "y", Feb: "f", Mar: "m", Apr: "a", May: "m", Jun: "y", Jul: "y", Aug: "a", Sep: "s", Oct: "ɔ", Nov: "n", Dec: "d"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "sánzá ya yambo", Feb: "sánzá ya míbalé", Mar: "sánzá ya mísáto", Apr: "sánzá ya mínei", May: "sánzá ya mítáno", Jun: "sánzá ya motóbá", Jul: "sánzá ya nsambo", Aug: "sánzá ya mwambe", Sep: "sánzá ya libwa", Oct: "sánzá ya zómi", Nov: "sánzá ya zómi na mɔ̌kɔ́", Dec: "sánzá ya zómi na míbalé"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "eye", Mon: "ybo", Tue: "mbl", Wed: "mst", Thu: "min", Fri: "mtn", Sat: "mps"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "e", Mon: "y", Tue: "m", Wed: "m", Thu: "m", Fri: "m", Sat: "p"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "eyenga", Mon: "mokɔlɔ mwa yambo", Tue: "mokɔlɔ mwa míbalé", Wed: "mokɔlɔ mwa mísáto", Thu: "mokɔlɔ ya mínéi", Fri: "mokɔlɔ ya mítáno", Sat: "mpɔ́sɔ"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "ntɔ́ngɔ́", PM: "mpókwa"},
			},
		},
	}
}
