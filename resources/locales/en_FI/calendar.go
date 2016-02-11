package en_FI

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{},
			Time:     ut.CalendarDateFormat{Full: "H.mm.ss zzzz", Long: "H.mm.ss z", Medium: "H.mm.ss", Short: "H.mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{},
	}
}
