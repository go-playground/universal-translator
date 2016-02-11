package fr_CH

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "", Medium: "", Short: "dd.MM.yy"},
			Time:     ut.CalendarDateFormat{Full: "HH.mm:ss 'h' zzzz", Long: "", Medium: "", Short: ""},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{},
	}
}
