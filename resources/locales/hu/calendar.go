package hu

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "y. MMMM d., EEEE", Long: "y. MMMM d.", Medium: "y. MMM d.", Short: "y. MM. dd."},
			Time:     ut.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "jan.", Feb: "febr.", Mar: "márc.", Apr: "ápr.", May: "máj.", Jun: "jún.", Jul: "júl.", Aug: "aug.", Sep: "szept.", Oct: "okt.", Nov: "nov.", Dec: "dec."},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "Á", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "Sz", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "január", Feb: "február", Mar: "március", Apr: "április", May: "május", Jun: "június", Jul: "július", Aug: "augusztus", Sep: "szeptember", Oct: "október", Nov: "november", Dec: "december"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "V", Mon: "H", Tue: "K", Wed: "Sze", Thu: "Cs", Fri: "P", Sat: "Szo"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "V", Mon: "H", Tue: "K", Wed: "Sz", Thu: "Cs", Fri: "P", Sat: "Sz"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "V", Mon: "H", Tue: "K", Wed: "Sze", Thu: "Cs", Fri: "P", Sat: "Szo"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "vasárnap", Mon: "hétfő", Tue: "kedd", Wed: "szerda", Thu: "csütörtök", Fri: "péntek", Sat: "szombat"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "de.", PM: "du."},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "de.", PM: "du."},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "de.", PM: "du."},
			},
		},
	}
}
