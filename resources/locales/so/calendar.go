package so

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, MMMM dd, y", Long: "dd MMMM y", Medium: "dd-MMM-y", Short: "dd/MM/yy"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Kob", Feb: "Lab", Mar: "Sad", Apr: "Afr", May: "Sha", Jun: "Lix", Jul: "Tod", Aug: "Sid", Sep: "Sag", Oct: "Tob", Nov: "KIT", Dec: "LIT"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "K", Feb: "L", Mar: "S", Apr: "A", May: "S", Jun: "L", Jul: "T", Aug: "S", Sep: "S", Oct: "T", Nov: "K", Dec: "L"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Bisha Koobaad", Feb: "Bisha Labaad", Mar: "Bisha Saddexaad", Apr: "Bisha Afraad", May: "Bisha Shanaad", Jun: "Bisha Lixaad", Jul: "Bisha Todobaad", Aug: "Bisha Sideedaad", Sep: "Bisha Sagaalaad", Oct: "Bisha Tobnaad", Nov: "Bisha Kow iyo Tobnaad", Dec: "Bisha Laba iyo Tobnaad"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Axd", Mon: "Isn", Tue: "Tal", Wed: "Arb", Thu: "Kha", Fri: "Jim", Sat: "Sab"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "A", Mon: "I", Tue: "T", Wed: "A", Thu: "K", Fri: "J", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Axad", Mon: "Isniin", Tue: "Talaado", Wed: "Arbaco", Thu: "Khamiis", Fri: "Jimco", Sat: "Sabti"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "sn.", PM: "gn."},
			},
		},
	}
}
