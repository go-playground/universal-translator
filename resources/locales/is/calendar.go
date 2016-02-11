package is

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM y", Short: "d.M.y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} 'kl'. {0}", Long: "{1} 'kl'. {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "jan.", Feb: "feb.", Mar: "mar.", Apr: "apr.", May: "maí", Jun: "jún.", Jul: "júl.", Aug: "ágú.", Sep: "sep.", Oct: "okt.", Nov: "nóv.", Dec: "des."},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "Á", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "janúar", Feb: "febrúar", Mar: "mars", Apr: "apríl", May: "maí", Jun: "júní", Jul: "júlí", Aug: "ágúst", Sep: "september", Oct: "október", Nov: "nóvember", Dec: "desember"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "sun.", Mon: "mán.", Tue: "þri.", Wed: "mið.", Thu: "fim.", Fri: "fös.", Sat: "lau."},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "Þ", Wed: "M", Thu: "F", Fri: "F", Sat: "L"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "su.", Mon: "má.", Tue: "þr.", Wed: "mi.", Thu: "fi.", Fri: "fö.", Sat: "la."},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "sunnudagur", Mon: "mánudagur", Tue: "þriðjudagur", Wed: "miðvikudagur", Thu: "fimmtudagur", Fri: "föstudagur", Sat: "laugardagur"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "f.h.", PM: "e.h."},
			},
		},
	}
}
