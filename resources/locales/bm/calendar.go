package bm

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "zan", Feb: "feb", Mar: "mar", Apr: "awi", May: "mɛ", Jun: "zuw", Jul: "zul", Aug: "uti", Sep: "sɛt", Oct: "ɔku", Nov: "now", Dec: "des"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "Z", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "Z", Jul: "Z", Aug: "U", Sep: "S", Oct: "Ɔ", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "zanwuye", Feb: "feburuye", Mar: "marisi", Apr: "awirili", May: "mɛ", Jun: "zuwɛn", Jul: "zuluye", Aug: "uti", Sep: "sɛtanburu", Oct: "ɔkutɔburu", Nov: "nowanburu", Dec: "desanburu"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "kar", Mon: "ntɛ", Tue: "tar", Wed: "ara", Thu: "ala", Fri: "jum", Sat: "sib"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "K", Mon: "N", Tue: "T", Wed: "A", Thu: "A", Fri: "J", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "kari", Mon: "ntɛnɛ", Tue: "tarata", Wed: "araba", Thu: "alamisa", Fri: "juma", Sat: "sibiri"},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
