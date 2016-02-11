package ar_MA

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "يناير", Feb: "فبراير", Mar: "مارس", Apr: "أبريل", May: "ماي", Jun: "يونيو", Jul: "يوليوز", Aug: "غشت", Sep: "شتنبر", Oct: "أكتوبر", Nov: "نونبر", Dec: "دجنبر"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "ي", Feb: "ف", Mar: "م", Apr: "أ", May: "م", Jun: "ن", Jul: "ل", Aug: "غ", Sep: "ش", Oct: "ك", Nov: "ب", Dec: "د"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "يناير", Feb: "فبراير", Mar: "مارس", Apr: "أبريل", May: "ماي", Jun: "يونيو", Jul: "يوليوز", Aug: "غشت", Sep: "شتنبر", Oct: "أكتوبر", Nov: "نونبر", Dec: "دجنبر"},
			},
			Days:    ut.CalendarDayFormatNames{},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
