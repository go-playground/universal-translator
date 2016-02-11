package it_CH

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "", Medium: "d-MMM-y", Short: "dd.MM.yy"},
			Time:     ut.CalendarDateFormat{},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "gennaio", Feb: "febbraio", Mar: "marzo", Apr: "aprile", May: "maggio", Jun: "giugno", Jul: "luglio", Aug: "agosto", Sep: "settembre", Oct: "ottobre", Nov: "novembre", Dec: "dicembre"},
			},
			Days:    ut.CalendarDayFormatNames{},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
