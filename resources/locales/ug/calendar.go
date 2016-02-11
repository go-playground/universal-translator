package ug

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE، MMMM d، y", Long: "MMMM d، y", Medium: "MMM d، y", Short: "M/d/yy"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}، {0}", Short: "{1}، {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "يانۋار", Feb: "فېۋرال", Mar: "مارت", Apr: "ئاپرېل", May: "ماي", Jun: "ئىيۇن", Jul: "ئىيۇل", Aug: "ئاۋغۇست", Sep: "سېنتەبىر", Oct: "ئۆكتەبىر", Nov: "نويابىر", Dec: "دېكابىر"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "يانۋار", Feb: "فېۋرال", Mar: "مارت", Apr: "ئاپرېل", May: "ماي", Jun: "ئىيۇن", Jul: "ئىيۇل", Aug: "ئاۋغۇست", Sep: "سېنتەبىر", Oct: "ئۆكتەبىر", Nov: "بويابىر", Dec: "دېكابىر"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "يە", Mon: "دۈ", Tue: "سە", Wed: "چا", Thu: "پە", Fri: "چۈ", Sat: "شە"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "ي", Mon: "د", Tue: "س", Wed: "چ", Thu: "پ", Fri: "ج", Sat: "ش"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "ي", Mon: "د", Tue: "س", Wed: "چ", Thu: "پ", Fri: "ج", Sat: "ش"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "يەكشەنبە", Mon: "دۈشەنبە", Tue: "سەيشەنبە", Wed: "چارشەنبە", Thu: "پەيشەنبە", Fri: "جۈمە", Sat: "شەنبە"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "چۈشتىن بۇرۇن", PM: "چۈشتىن كېيىن"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "چۈشتىن بۇرۇن", PM: "چۈشتىن كېيىن"},
			},
		},
	}
}
