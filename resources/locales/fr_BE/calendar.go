package fr_BE

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: "d/MM/yy"},
			Time:     ut.CalendarDateFormat{Full: "H 'h' mm 'min' ss 's' zzzz", Long: "", Medium: "", Short: ""},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{},
	}
}
