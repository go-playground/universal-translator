package nd

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Zib", Feb: "Nhlo", Mar: "Mbi", Apr: "Mab", May: "Nkw", Jun: "Nhla", Jul: "Ntu", Aug: "Ncw", Sep: "Mpan", Oct: "Mfu", Nov: "Lwe", Dec: "Mpal"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "Z", Feb: "N", Mar: "M", Apr: "M", May: "N", Jun: "N", Jul: "N", Aug: "N", Sep: "M", Oct: "M", Nov: "L", Dec: "M"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Zibandlela", Feb: "Nhlolanja", Mar: "Mbimbitho", Apr: "Mabasa", May: "Nkwenkwezi", Jun: "Nhlangula", Jul: "Ntulikazi", Aug: "Ncwabakazi", Sep: "Mpandula", Oct: "Mfumfu", Nov: "Lwezi", Dec: "Mpalakazi"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Son", Mon: "Mvu", Tue: "Sib", Wed: "Sit", Thu: "Sin", Fri: "Sih", Sat: "Mgq"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "S", Wed: "S", Thu: "S", Fri: "S", Sat: "M"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Sonto", Mon: "Mvulo", Tue: "Sibili", Wed: "Sithathu", Thu: "Sine", Fri: "Sihlanu", Sat: "Mgqibelo"},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
