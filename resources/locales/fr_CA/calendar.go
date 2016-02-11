package fr_CA

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: "yy-MM-dd"},
			Time:     ut.CalendarDateFormat{},
			DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "{1} {0}", Short: ""},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{},
			Days:   ut.CalendarDayFormatNames{},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{},
			},
		},
	}
}
