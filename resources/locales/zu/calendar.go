package zu

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mas", Apr: "Eph", May: "Mey", Jun: "Jun", Jul: "Jul", Aug: "Aga", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Dis"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Januwari", Feb: "Februwari", Mar: "Mashi", Apr: "Ephreli", May: "Meyi", Jun: "Juni", Jul: "Julayi", Aug: "Agasti", Sep: "Septhemba", Oct: "Okthoba", Nov: "Novemba", Dec: "Disemba"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Son", Mon: "Mso", Tue: "Bil", Wed: "Tha", Thu: "Sin", Fri: "Hla", Sat: "Mgq"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "B", Wed: "T", Thu: "S", Fri: "H", Sat: "M"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "Son", Mon: "Mso", Tue: "Bil", Wed: "Tha", Thu: "Sin", Fri: "Hla", Sat: "Mgq"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "ISonto", Mon: "UMsombuluko", Tue: "ULwesibili", Wed: "ULwesithathu", Thu: "ULwesine", Fri: "ULwesihlanu", Sat: "UMgqibelo"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			},
		},
	}
}
