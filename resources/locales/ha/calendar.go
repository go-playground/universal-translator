package ha

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "d/M/yy"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Fab", Mar: "Mar", Apr: "Afi", May: "May", Jun: "Yun", Jul: "Yul", Aug: "Agu", Sep: "Sat", Oct: "Okt", Nov: "Nuw", Dec: "Dis"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "Y", Jul: "Y", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Janairu", Feb: "Faburairu", Mar: "Maris", Apr: "Afirilu", May: "Mayu", Jun: "Yuni", Jul: "Yuli", Aug: "Agusta", Sep: "Satumba", Oct: "Oktoba", Nov: "Nuwamba", Dec: "Disamba"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Lh", Mon: "Li", Tue: "Ta", Wed: "Lr", Thu: "Al", Fri: "Ju", Sat: "As"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "L", Mon: "L", Tue: "T", Wed: "L", Thu: "A", Fri: "J", Sat: "A"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Lahadi", Mon: "Litinin", Tue: "Talata", Wed: "Laraba", Thu: "Alhamis", Fri: "Jummaʼa", Sat: "Asabar"},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
