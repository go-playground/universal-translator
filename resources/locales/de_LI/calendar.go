package de_LI

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{},
			Days:   ut.CalendarDayFormatNames{},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "vm.", PM: "nm."},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{},
			},
		},
	}
}
