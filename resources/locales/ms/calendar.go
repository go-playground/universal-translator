package ms

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/MM/yy"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mac", Apr: "Apr", May: "Mei", Jun: "Jun", Jul: "Jul", Aug: "Ogo", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Dis"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "O", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Januari", Feb: "Februari", Mar: "Mac", Apr: "April", May: "Mei", Jun: "Jun", Jul: "Julai", Aug: "Ogos", Sep: "September", Oct: "Oktober", Nov: "November", Dec: "Disember"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Ahd", Mon: "Isn", Tue: "Sel", Wed: "Rab", Thu: "Kha", Fri: "Jum", Sat: "Sab"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "A", Mon: "I", Tue: "S", Wed: "R", Thu: "K", Fri: "J", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "Ah", Mon: "Is", Tue: "Se", Wed: "Ra", Thu: "Kh", Fri: "Ju", Sat: "Sa"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Ahad", Mon: "Isnin", Tue: "Selasa", Wed: "Rabu", Thu: "Khamis", Fri: "Jumaat", Sat: "Sabtu"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "PG", PM: "PTG"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "PG", PM: "PTG"},
			},
		},
	}
}
