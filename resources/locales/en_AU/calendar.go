package en_AU

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "", Long: "dMMMM,y", Medium: "dMMM,y", Short: "d/M/yy"},
			Time:     ut.CalendarDateFormat{},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "May", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "Jan.", Feb: "Feb.", Mar: "Mar.", Apr: "Apr.", May: "May", Jun: "Jun.", Jul: "Jul.", Aug: "Aug.", Sep: "Sep.", Oct: "Oct.", Nov: "Nov.", Dec: "Dec."},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Sun.", Mon: "Mon.", Tue: "Tue.", Wed: "Wed.", Thu: "Thu.", Fri: "Fri.", Sat: "Sat."},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "Su.", Mon: "M.", Tue: "Tu.", Wed: "W.", Thu: "Th.", Fri: "F.", Sat: "Sa."},
				Short:       ut.CalendarDayFormatNameValue{Sun: "Su.", Mon: "Mon.", Tue: "Tu.", Wed: "Wed.", Thu: "Th.", Fri: "Fri.", Sat: "Sat."},
				Wide:        ut.CalendarDayFormatNameValue{},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "am", PM: "pm"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "am", PM: "pm"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "am", PM: "pm"},
			},
		},
	}
}
