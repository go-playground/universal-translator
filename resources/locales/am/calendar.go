package am

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE ፣d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "ጃንዩ", Feb: "ፌብሩ", Mar: "ማርች", Apr: "ኤፕሪ", May: "ሜይ", Jun: "ጁን", Jul: "ጁላይ", Aug: "ኦገስ", Sep: "ሴፕቴ", Oct: "ኦክቶ", Nov: "ኖቬም", Dec: "ዲሴም"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "ጃ", Feb: "ፌ", Mar: "ማ", Apr: "ኤ", May: "ሜ", Jun: "ጁ", Jul: "ጁ", Aug: "ኦ", Sep: "ሴ", Oct: "ኦ", Nov: "ኖ", Dec: "ዲ"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "ጃንዩወሪ", Feb: "ፌብሩወሪ", Mar: "ማርች", Apr: "ኤፕሪል", May: "ሜይ", Jun: "ጁን", Jul: "ጁላይ", Aug: "ኦገስት", Sep: "ሴፕቴምበር", Oct: "ኦክቶበር", Nov: "ኖቬምበር", Dec: "ዲሴምበር"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "እሑድ", Mon: "ሰኞ", Tue: "ማክሰ", Wed: "ረቡዕ", Thu: "ሐሙስ", Fri: "ዓርብ", Sat: "ቅዳሜ"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "እ", Mon: "ሰ", Tue: "ማ", Wed: "ረ", Thu: "ሐ", Fri: "ዓ", Sat: "ቅ"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "እ", Mon: "ሰ", Tue: "ማ", Wed: "ረ", Thu: "ሐ", Fri: "ዓ", Sat: "ቅ"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "እሑድ", Mon: "ሰኞ", Tue: "ማክሰኞ", Wed: "ረቡዕ", Thu: "ሐሙስ", Fri: "ዓርብ", Sat: "ቅዳሜ"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "ጥዋት", PM: "ከሰዓት"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "ጠ", PM: "ከ"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "ጥዋት", PM: "ከሰዓት"},
			},
		},
	}
}
