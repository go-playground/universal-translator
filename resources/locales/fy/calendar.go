package fy

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "jan", Feb: "feb", Mar: "mrt", Apr: "apr", May: "mai", Jun: "jun", Jul: "jul", Aug: "aug", Sep: "sep", Oct: "okt", Nov: "nov", Dec: "des"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "jannewaris", Feb: "febrewaris", Mar: "maart", Apr: "april", May: "maaie", Jun: "juny", Jul: "july", Aug: "augustus", Sep: "septimber", Oct: "oktober", Nov: "novimber", Dec: "desimber"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "si", Mon: "mo", Tue: "ti", Wed: "wo", Thu: "to", Fri: "fr", Sat: "so"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "Z", Mon: "M", Tue: "D", Wed: "W", Thu: "D", Fri: "V", Sat: "Z"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "si", Mon: "mo", Tue: "ti", Wed: "wo", Thu: "to", Fri: "fr", Sat: "so"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "snein", Mon: "moandei", Tue: "tiisdei", Wed: "woansdei", Thu: "tongersdei", Fri: "freed", Sat: "sneon"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "foarmiddei", PM: "p.m."},
			},
		},
	}
}
