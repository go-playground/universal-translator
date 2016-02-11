package bem

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mac", Apr: "Epr", May: "Mei", Jun: "Jun", Jul: "Jul", Aug: "Oga", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Dis"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "E", May: "M", Jun: "J", Jul: "J", Aug: "O", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Januari", Feb: "Februari", Mar: "Machi", Apr: "Epreo", May: "Mei", Jun: "Juni", Jul: "Julai", Aug: "Ogasti", Sep: "Septemba", Oct: "Oktoba", Nov: "Novemba", Dec: "Disemba"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{},
				Narrow:      ut.CalendarDayFormatNameValue{},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Pa Mulungu", Mon: "Palichimo", Tue: "Palichibuli", Wed: "Palichitatu", Thu: "Palichine", Fri: "Palichisano", Sat: "Pachibelushi"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "uluchelo", PM: "akasuba"},
			},
		},
	}
}
