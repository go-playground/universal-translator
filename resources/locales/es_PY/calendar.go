package es_PY

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "ene.", Feb: "feb.", Mar: "mar.", Apr: "abr.", May: "may.", Jun: "jun.", Jul: "jul.", Aug: "ago.", Sep: "sept.", Oct: "oct.", Nov: "nov.", Dec: "dic."},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{},
				Narrow:      ut.CalendarDayFormatNameValue{},
				Short:       ut.CalendarDayFormatNameValue{Sun: "Do", Mon: "Lu", Tue: "Ma", Wed: "Mi", Thu: "Ju", Fri: "Vi", Sat: "Sa"},
				Wide:        ut.CalendarDayFormatNameValue{},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
