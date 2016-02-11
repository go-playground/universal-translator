package ti_ER

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE፡ dd MMMM መዓልቲ y G", Long: "", Medium: "", Short: ""},
			Time:     ut.CalendarDateFormat{},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "ጥሪ", Feb: "ለካቲ", Mar: "መጋቢ", Apr: "ሚያዝ", May: "ግንቦ", Jun: "ሰነ", Jul: "ሓምለ", Aug: "ነሓሰ", Sep: "መስከ", Oct: "ጥቅም", Nov: "ሕዳር", Dec: "ታሕሳ"},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "ጥሪ", Feb: "ለካቲት", Mar: "መጋቢት", Apr: "ሚያዝያ", May: "ግንቦት", Jun: "ሰነ", Jul: "ሓምለ", Aug: "ነሓሰ", Sep: "መስከረም", Oct: "ጥቅምቲ", Nov: "ሕዳር", Dec: "ታሕሳስ"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{},
				Narrow:      ut.CalendarDayFormatNameValue{},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "ሰንበት", Mon: "ሰኑይ", Tue: "ሰሉስ", Wed: "ረቡዕ", Thu: "ሓሙስ", Fri: "ዓርቢ", Sat: "ቀዳም"},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
