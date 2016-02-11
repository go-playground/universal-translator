package ar_MR

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "يناير", Feb: "فبراير", Mar: "مارس", Apr: "إبريل", May: "مايو", Jun: "يونيو", Jul: "يوليو", Aug: "أغشت", Sep: "شتمبر", Oct: "أكتوبر", Nov: "نوفمبر", Dec: "دجمبر"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "إ", May: "", Jun: "", Jul: "", Aug: "", Sep: "ش", Oct: "", Nov: "", Dec: ""},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "يناير", Feb: "فبراير", Mar: "مارس", Apr: "إبريل", May: "مايو", Jun: "يونيو", Jul: "يوليو", Aug: "أغشت", Sep: "شتمبر", Oct: "أكتوبر", Nov: "نوفمبر", Dec: "دجمبر"},
			},
			Days:    ut.CalendarDayFormatNames{},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
