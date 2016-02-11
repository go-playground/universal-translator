package lv

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, y. 'gada' d. MMMM", Long: "y. 'gada' d. MMMM", Medium: "y. 'gada' d. MMM", Short: "dd.MM.yy"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Janv.", Feb: "Febr.", Mar: "Marts", Apr: "Apr.", May: "Maijs", Jun: "Jūn.", Jul: "Jūl.", Aug: "Aug.", Sep: "Sept.", Oct: "Okt.", Nov: "Nov.", Dec: "Dec."},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Janvāris", Feb: "Februāris", Mar: "Marts", Apr: "Aprīlis", May: "Maijs", Jun: "Jūnijs", Jul: "Jūlijs", Aug: "Augusts", Sep: "Septembris", Oct: "Oktobris", Nov: "Novembris", Dec: "Decembris"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Sv", Mon: "Pr", Tue: "Ot", Wed: "Tr", Thu: "Ce", Fri: "Pk", Sat: "Se"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "S", Mon: "P", Tue: "O", Wed: "T", Thu: "C", Fri: "P", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "Sv", Mon: "Pr", Tue: "Ot", Wed: "Tr", Thu: "Ce", Fri: "Pk", Sat: "Se"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Svētdiena", Mon: "Pirmdiena", Tue: "Otrdiena", Wed: "Trešdiena", Thu: "Ceturtdiena", Fri: "Piektdiena", Sat: "Sestdiena"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "priekšp.", PM: "pēcpusd."},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "priekšp.", PM: "pēcp."},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "priekšpusdienā", PM: "pēcpusdiena"},
			},
		},
	}
}
