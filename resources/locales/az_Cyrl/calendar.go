package az_Cyrl

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d, MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "dd.MM.yy"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "јанвар", Feb: "феврал", Mar: "март", Apr: "апрел", May: "май", Jun: "ијун", Jul: "ијул", Aug: "август", Sep: "сентјабр", Oct: "октјабр", Nov: "нојабр", Dec: "декабр"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "7", Mon: "1", Tue: "2", Wed: "3", Thu: "4", Fri: "5", Sat: "6"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "базар", Mon: "базар ертәси", Tue: "чәршәнбә ахшамы", Wed: "чәршәнбә", Thu: "ҹүмә ахшамы", Fri: "ҹүмә", Sat: "шәнбә"},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
