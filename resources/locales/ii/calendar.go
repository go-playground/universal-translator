package ii

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "ꋍꆪ", Feb: "ꑍꆪ", Mar: "ꌕꆪ", Apr: "ꇖꆪ", May: "ꉬꆪ", Jun: "ꃘꆪ", Jul: "ꏃꆪ", Aug: "ꉆꆪ", Sep: "ꈬꆪ", Oct: "ꊰꆪ", Nov: "ꊰꊪꆪ", Dec: "ꊰꑋꆪ"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "ꑭꆏ", Mon: "ꆏꋍ", Tue: "ꆏꑍ", Wed: "ꆏꌕ", Thu: "ꆏꇖ", Fri: "ꆏꉬ", Sat: "ꆏꃘ"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "ꆏ", Mon: "ꋍ", Tue: "ꑍ", Wed: "ꌕ", Thu: "ꇖ", Fri: "ꉬ", Sat: "ꃘ"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "ꑭꆏꑍ", Mon: "ꆏꊂꋍ", Tue: "ꆏꊂꑍ", Wed: "ꆏꊂꌕ", Thu: "ꆏꊂꇖ", Fri: "ꆏꊂꉬ", Sat: "ꆏꊂꃘ"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "ꎸꄑ", PM: "ꁯꋒ"},
			},
		},
	}
}
