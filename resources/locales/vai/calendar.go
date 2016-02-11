package vai

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "ꖨꕪꖃ ꔞꕮ", Feb: "ꕒꕡꖝꖕ", Mar: "ꕾꖺ", Apr: "ꖢꖕ", May: "ꖑꕱ", Jun: "6", Jul: "7", Aug: "ꗛꔕ", Sep: "ꕢꕌ", Oct: "ꕭꖃ", Nov: "ꔞꘋꕔꕿ ꕸꖃꗏ", Dec: "ꖨꕪꕱ ꗏꕮ"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{},
				Narrow:      ut.CalendarDayFormatNameValue{},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "ꕞꕌꔵ", Mon: "ꗳꗡꘉ", Tue: "ꕚꕞꕚ", Wed: "ꕉꕞꕒ", Thu: "ꕉꔤꕆꕢ", Fri: "ꕉꔤꕀꕮ", Sat: "ꔻꔬꔳ"},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
