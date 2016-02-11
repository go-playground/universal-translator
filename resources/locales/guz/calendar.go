package guz

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Can", Feb: "Feb", Mar: "Mac", Apr: "Apr", May: "Mei", Jun: "Jun", Jul: "Cul", Aug: "Agt", Sep: "Sep", Oct: "Okt", Nov: "Nob", Dec: "Dis"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "C", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "C", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Chanuari", Feb: "Feburari", Mar: "Machi", Apr: "Apiriri", May: "Mei", Jun: "Juni", Jul: "Chulai", Aug: "Agosti", Sep: "Septemba", Oct: "Okitoba", Nov: "Nobemba", Dec: "Disemba"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Cpr", Mon: "Ctt", Tue: "Cmn", Wed: "Cmt", Thu: "Ars", Fri: "Icm", Sat: "Est"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "C", Mon: "C", Tue: "C", Wed: "C", Thu: "A", Fri: "I", Sat: "E"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Chumapiri", Mon: "Chumatato", Tue: "Chumaine", Wed: "Chumatano", Thu: "Aramisi", Fri: "Ichuma", Sat: "Esabato"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "Ma/Mo", PM: "Mambia/Mog"},
			},
		},
	}
}
