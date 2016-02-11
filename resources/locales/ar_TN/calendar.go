package ar_TN

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "جانفي", Feb: "فيفري", Mar: "مارس", Apr: "أفريل", May: "ماي", Jun: "جوان", Jul: "جويلية", Aug: "أوت", Sep: "سبتمبر", Oct: "أكتوبر", Nov: "نوفمبر", Dec: "ديسمبر"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "ج", Feb: "ف", Mar: "م", Apr: "أ", May: "م", Jun: "ج", Jul: "ج", Aug: "أ", Sep: "س", Oct: "أ", Nov: "ن", Dec: "د"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "جانفي", Feb: "فيفري", Mar: "مارس", Apr: "أفريل", May: "ماي", Jun: "جوان", Jul: "جويلية", Aug: "أوت", Sep: "سبتمبر", Oct: "أكتوبر", Nov: "نوفمبر", Dec: "ديسمبر"},
			},
			Days:    ut.CalendarDayFormatNames{},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
