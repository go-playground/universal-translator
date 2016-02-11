package ro_MD

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Dum", Mon: "Lun", Tue: "Mar", Wed: "Mie", Thu: "Joi", Fri: "Vin", Sat: "Sâm"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "Ma", Wed: "Mi", Thu: "J", Fri: "V", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "Du", Mon: "Lu", Tue: "Ma", Wed: "Mi", Thu: "Jo", Fri: "Vi", Sat: "Sâ"},
				Wide:        ut.CalendarDayFormatNameValue{},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
