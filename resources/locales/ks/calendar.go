package ks

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "ج", Feb: "ف", Mar: "م", Apr: "ا", May: "م", Jun: "ج", Jul: "ج", Aug: "ا", Sep: "س", Oct: "س", Nov: "ا", Dec: "ن"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "جنؤری", Feb: "فرؤری", Mar: "مارٕچ", Apr: "اپریل", May: "میٔ", Jun: "جوٗن", Jul: "جوٗلایی", Aug: "اگست", Sep: "ستمبر", Oct: "اکتوٗبر", Nov: "نومبر", Dec: "دسمبر"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "آتھوار", Mon: "ژٔنٛدٕروار", Tue: "بوٚموار", Wed: "بودوار", Thu: "برٛٮ۪سوار", Fri: "جُمہ", Sat: "بٹوار"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "ا", Mon: "ژ", Tue: "ب", Wed: "ب", Thu: "ب", Fri: "ج", Sat: "ب"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "اَتھوار", Mon: "ژٔنٛدرٕروار", Tue: "بوٚموار", Wed: "بودوار", Thu: "برٛٮ۪سوار", Fri: "جُمہ", Sat: "بٹوار"},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
