package pa

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/yy"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "ਜਨ", Feb: "ਫ਼ਰ", Mar: "ਮਾਰਚ", Apr: "ਅਪ੍ਰੈ", May: "ਮਈ", Jun: "ਜੂਨ", Jul: "ਜੁਲਾ", Aug: "ਅਗ", Sep: "ਸਤੰ", Oct: "ਅਕਤੂ", Nov: "ਨਵੰ", Dec: "ਦਸੰ"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "ਜ", Feb: "ਫ਼", Mar: "ਮਾ", Apr: "ਅ", May: "ਮ", Jun: "ਜੂ", Jul: "ਜੁ", Aug: "ਅ", Sep: "ਸ", Oct: "ਅ", Nov: "ਨ", Dec: "ਦ"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "ਜਨਵਰੀ", Feb: "ਫ਼ਰਵਰੀ", Mar: "ਮਾਰਚ", Apr: "ਅਪ੍ਰੈਲ", May: "ਮਈ", Jun: "ਜੂਨ", Jul: "ਜੁਲਾਈ", Aug: "ਅਗਸਤ", Sep: "ਸਤੰਬਰ", Oct: "ਅਕਤੂਬਰ", Nov: "ਨਵੰਬਰ", Dec: "ਦਸੰਬਰ"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "ਐਤ", Mon: "ਸੋਮ", Tue: "ਮੰਗਲ", Wed: "ਬੁੱਧ", Thu: "ਵੀਰ", Fri: "ਸ਼ੁੱਕਰ", Sat: "ਸ਼ਨਿੱਚਰ"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "ਐ", Mon: "ਸੋ", Tue: "ਮੰ", Wed: "ਬੁੱ", Thu: "ਵੀ", Fri: "ਸ਼ੁੱ", Sat: "ਸ਼"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "ਐਤ", Mon: "ਸੋਮ", Tue: "ਮੰਗ", Wed: "ਬੁੱਧ", Thu: "ਵੀਰ", Fri: "ਸ਼ੁੱਕ", Sat: "ਸ਼ਨਿੱ"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "ਐਤਵਾਰ", Mon: "ਸੋਮਵਾਰ", Tue: "ਮੰਗਲਵਾਰ", Wed: "ਬੁੱਧਵਾਰ", Thu: "ਵੀਰਵਾਰ", Fri: "ਸ਼ੁੱਕਰਵਾਰ", Sat: "ਸ਼ਨਿੱਚਰਵਾਰ"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "ਪੂ.ਦੁ.", PM: "ਬਾ.ਦੁ."},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "ਪੂ.ਦੁ.", PM: "ਬਾ.ਦੁ."},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "ਪੂ.ਦੁ.", PM: "ਬਾ.ਦੁ."},
			},
		},
	}
}
