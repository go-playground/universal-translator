package fa_AF

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "جنو", Feb: "", Mar: "", Apr: "", May: "می", Jun: "", Jul: "جول", Aug: "", Sep: "", Oct: "", Nov: "", Dec: "دسم"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "ج", Feb: "ف", Mar: "م", Apr: "ا", May: "م", Jun: "ج", Jul: "ج", Aug: "ا", Sep: "س", Oct: "ا", Nov: "ن", Dec: "د"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "جنوری", Feb: "فبروری", Mar: "مارچ", Apr: "اپریل", May: "می", Jun: "جون", Jul: "جولای", Aug: "اگست", Sep: "سپتمبر", Oct: "اکتوبر", Nov: "نومبر", Dec: "دسمبر"},
			},
			Days:    ut.CalendarDayFormatNames{},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
