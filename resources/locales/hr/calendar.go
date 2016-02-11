package hr

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d. MMMM y.", Long: "d. MMMM y.", Medium: "d. MMM y.", Short: "dd.MM.y."},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} 'u' {0}", Long: "{1} 'u' {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "sij", Feb: "velj", Mar: "ožu", Apr: "tra", May: "svi", Jun: "lip", Jul: "srp", Aug: "kol", Sep: "ruj", Oct: "lis", Nov: "stu", Dec: "pro"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "1.", Feb: "2.", Mar: "3.", Apr: "4.", May: "5.", Jun: "6.", Jul: "7.", Aug: "8.", Sep: "9.", Oct: "10.", Nov: "11.", Dec: "12."},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "siječanj", Feb: "veljača", Mar: "ožujak", Apr: "travanj", May: "svibanj", Jun: "lipanj", Jul: "srpanj", Aug: "kolovoz", Sep: "rujan", Oct: "listopad", Nov: "studeni", Dec: "prosinac"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "ned", Mon: "pon", Tue: "uto", Wed: "sri", Thu: "čet", Fri: "pet", Sat: "sub"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "n", Mon: "p", Tue: "u", Wed: "s", Thu: "č", Fri: "p", Sat: "s"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "ned", Mon: "pon", Tue: "uto", Wed: "sri", Thu: "čet", Fri: "pet", Sat: "sub"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "nedjelja", Mon: "ponedjeljak", Tue: "utorak", Wed: "srijeda", Thu: "četvrtak", Fri: "petak", Sat: "subota"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			},
		},
	}
}
