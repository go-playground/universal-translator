package ckb

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "", Long: "dی MMMMی y", Medium: "", Short: ""},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "کانوونی دووەم", Feb: "شوبات", Mar: "ئازار", Apr: "نیسان", May: "ئایار", Jun: "حوزەیران", Jul: "تەمووز", Aug: "ئاب", Sep: "ئەیلوول", Oct: "تشرینی یەکەم", Nov: "تشرینی دووەم", Dec: "کانونی یەکەم"},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "کانوونی دووەم", Feb: "شوبات", Mar: "ئازار", Apr: "نیسان", May: "ئایار", Jun: "حوزەیران", Jul: "تەمووز", Aug: "ئاب", Sep: "ئەیلوول", Oct: "تشرینی یەکەم", Nov: "تشرینی دووەم", Dec: "کانونی یەکەم"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "یەکشەممە", Mon: "دووشەممە", Tue: "سێشەممە", Wed: "چوارشەممە", Thu: "پێنجشەممە", Fri: "ھەینی", Sat: "شەممە"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "ی", Mon: "د", Tue: "س", Wed: "چ", Thu: "پ", Fri: "ھ", Sat: "ش"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "یەکشەممە", Mon: "دووشەممە", Tue: "سێشەممە", Wed: "چوارشەممە", Thu: "پێنجشەممە", Fri: "ھەینی", Sat: "شەممە"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "ب.ن", PM: "د.ن"},
			},
		},
	}
}
