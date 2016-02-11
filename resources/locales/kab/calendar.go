package kab

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Yen", Feb: "Fur", Mar: "Meɣ", Apr: "Yeb", May: "May", Jun: "Yun", Jul: "Yul", Aug: "Ɣuc", Sep: "Cte", Oct: "Tub", Nov: "Nun", Dec: "Duǧ"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "Y", Feb: "F", Mar: "M", Apr: "Y", May: "M", Jun: "Y", Jul: "Y", Aug: "Ɣ", Sep: "C", Oct: "T", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Yennayer", Feb: "Fuṛar", Mar: "Meɣres", Apr: "Yebrir", May: "Mayyu", Jun: "Yunyu", Jul: "Yulyu", Aug: "Ɣuct", Sep: "Ctembeṛ", Oct: "Tubeṛ", Nov: "Nunembeṛ", Dec: "Duǧembeṛ"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Yan", Mon: "San", Tue: "Kraḍ", Wed: "Kuẓ", Thu: "Sam", Fri: "Sḍis", Sat: "Say"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "Y", Mon: "S", Tue: "K", Wed: "K", Thu: "S", Fri: "S", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Yanass", Mon: "Sanass", Tue: "Kraḍass", Wed: "Kuẓass", Thu: "Samass", Fri: "Sḍisass", Sat: "Sayass"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "n tufat", PM: "n tmeddit"},
			},
		},
	}
}
