package yi

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, dטן MMMM y", Long: "dטן MMMM y", Medium: "dטן MMM y", Short: "dd/MM/yy"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}, {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "יאַנ", Feb: "פֿעב", Mar: "מערץ", Apr: "אַפּר", May: "מיי", Jun: "יוני", Jul: "יולי", Aug: "אויג", Sep: "סעפּ", Oct: "אקט", Nov: "נאוו", Dec: "דעצ"},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "יאַנואַר", Feb: "פֿעברואַר", Mar: "מערץ", Apr: "אַפּריל", May: "מיי", Jun: "יוני", Jul: "יולי", Aug: "אויגוסט", Sep: "סעפּטעמבער", Oct: "אקטאבער", Nov: "נאוועמבער", Dec: "דעצעמבער"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "זונטיק", Mon: "מאָנטיק", Tue: "דינסטיק", Wed: "מיטוואך", Thu: "דאנערשטיק", Fri: "פֿרײַטיק", Sat: "שבת"},
				Narrow:      ut.CalendarDayFormatNameValue{},
				Short:       ut.CalendarDayFormatNameValue{Sun: "זונטיק", Mon: "מאָנטיק", Tue: "דינסטיק", Wed: "מיטוואך", Thu: "דאנערשטיק", Fri: "פֿרײַטיק", Sat: "שבת"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "זונטיק", Mon: "מאָנטיק", Tue: "דינסטיק", Wed: "מיטוואך", Thu: "דאנערשטיק", Fri: "פֿרײַטיק", Sat: "שבת"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "פֿאַרמיטאָג", PM: "נאָכמיטאָג"},
			},
		},
	}
}
