package ar_LB

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "كانون الثاني", Feb: "شباط", Mar: "آذار", Apr: "نيسان", May: "أيار", Jun: "حزيران", Jul: "تموز", Aug: "آب", Sep: "أيلول", Oct: "تشرين الأول", Nov: "تشرين الثاني", Dec: "كانون الأول"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "ك", Feb: "ش", Mar: "آ", Apr: "ن", May: "أ", Jun: "ح", Jul: "ت", Aug: "آ", Sep: "أ", Oct: "ت", Nov: "ت", Dec: "ك"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "كانون الثاني", Feb: "شباط", Mar: "آذار", Apr: "نيسان", May: "أيار", Jun: "حزيران", Jul: "تموز", Aug: "آب", Sep: "أيلول", Oct: "تشرين الأول", Nov: "تشرين الثاني", Dec: "كانون الأول"},
			},
			Days:    ut.CalendarDayFormatNames{},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
