package om

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "dd MMMM y", Medium: "dd-MMM-y", Short: "dd/MM/yy"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Ama", Feb: "Gur", Mar: "Bit", Apr: "Elb", May: "Cam", Jun: "Wax", Jul: "Ado", Aug: "Hag", Sep: "Ful", Oct: "Onk", Nov: "Sad", Dec: "Mud"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Amajjii", Feb: "Guraandhala", Mar: "Bitooteessa", Apr: "Elba", May: "Caamsa", Jun: "Waxabajjii", Jul: "Adooleessa", Aug: "Hagayya", Sep: "Fuulbana", Oct: "Onkololeessa", Nov: "Sadaasa", Dec: "Muddee"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Dil", Mon: "Wix", Tue: "Qib", Wed: "Rob", Thu: "Kam", Fri: "Jim", Sat: "San"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "W", Thu: "T", Fri: "F", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Dilbata", Mon: "Wiixata", Tue: "Qibxata", Wed: "Roobii", Thu: "Kamiisa", Fri: "Jimaata", Sat: "Sanbata"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "WD", PM: "WB"},
			},
		},
	}
}
