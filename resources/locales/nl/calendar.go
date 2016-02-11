package nl

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd-MM-yy"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "jan.", Feb: "feb.", Mar: "mrt.", Apr: "apr.", May: "mei", Jun: "jun.", Jul: "jul.", Aug: "aug.", Sep: "sep.", Oct: "okt.", Nov: "nov.", Dec: "dec."},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "januari", Feb: "februari", Mar: "maart", Apr: "april", May: "mei", Jun: "juni", Jul: "juli", Aug: "augustus", Sep: "september", Oct: "oktober", Nov: "november", Dec: "december"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "zo", Mon: "ma", Tue: "di", Wed: "wo", Thu: "do", Fri: "vr", Sat: "za"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "Z", Mon: "M", Tue: "D", Wed: "W", Thu: "D", Fri: "V", Sat: "Z"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "zo", Mon: "ma", Tue: "di", Wed: "wo", Thu: "do", Fri: "vr", Sat: "za"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "zondag", Mon: "maandag", Tue: "dinsdag", Wed: "woensdag", Thu: "donderdag", Fri: "vrijdag", Sat: "zaterdag"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
			},
		},
	}
}
