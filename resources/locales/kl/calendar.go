package kl

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE dd MMMM y", Long: "dd MMMM y", Medium: "MMM dd, y", Short: "y-MM-dd"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "jan", Feb: "feb", Mar: "mar", Apr: "apr", May: "maj", Jun: "jun", Jul: "jul", Aug: "aug", Sep: "sep", Oct: "okt", Nov: "nov", Dec: "dec"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "januari", Feb: "februari", Mar: "martsi", Apr: "aprili", May: "maji", Jun: "juni", Jul: "juli", Aug: "augustusi", Sep: "septemberi", Oct: "oktoberi", Nov: "novemberi", Dec: "decemberi"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "sab", Mon: "ata", Tue: "mar", Wed: "pin", Thu: "sis", Fri: "tal", Sat: "arf"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "S", Mon: "A", Tue: "M", Wed: "P", Thu: "S", Fri: "T", Sat: "A"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "sab", Mon: "ata", Tue: "mar", Wed: "pin", Thu: "sis", Fri: "tal", Sat: "arf"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "sabaat", Mon: "ataasinngorneq", Tue: "marlunngorneq", Wed: "pingasunngorneq", Thu: "sisamanngorneq", Fri: "tallimanngorneq", Sat: "arfininngorneq"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "u.t.", PM: "u.k."},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "ulloqeqqata-tungaa", PM: "ulloqeqqata-kingorna"},
			},
		},
	}
}
