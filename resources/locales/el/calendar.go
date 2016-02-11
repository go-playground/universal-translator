package el

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/yy"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{Full: "{1} - {0}", Long: "{1} - {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Ιαν", Feb: "Φεβ", Mar: "Μάρ", Apr: "Απρ", May: "Μάι", Jun: "Ιούν", Jul: "Ιούλ", Aug: "Αύγ", Sep: "Σεπ", Oct: "Οκτ", Nov: "Νοέ", Dec: "Δεκ"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "Ι", Feb: "Φ", Mar: "Μ", Apr: "Α", May: "Μ", Jun: "Ι", Jul: "Ι", Aug: "Α", Sep: "Σ", Oct: "Ο", Nov: "Ν", Dec: "Δ"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Ιανουάριος", Feb: "Φεβρουάριος", Mar: "Μάρτιος", Apr: "Απρίλιος", May: "Μάιος", Jun: "Ιούνιος", Jul: "Ιούλιος", Aug: "Αύγουστος", Sep: "Σεπτέμβριος", Oct: "Οκτώβριος", Nov: "Νοέμβριος", Dec: "Δεκέμβριος"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Κυρ", Mon: "Δευ", Tue: "Τρί", Wed: "Τετ", Thu: "Πέμ", Fri: "Παρ", Sat: "Σάβ"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "Κ", Mon: "Δ", Tue: "Τ", Wed: "Τ", Thu: "Π", Fri: "Π", Sat: "Σ"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "Κυ", Mon: "Δε", Tue: "Τρ", Wed: "Τε", Thu: "Πέ", Fri: "Πα", Sat: "Σά"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Κυριακή", Mon: "Δευτέρα", Tue: "Τρίτη", Wed: "Τετάρτη", Thu: "Πέμπτη", Fri: "Παρασκευή", Sat: "Σάββατο"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "π.μ.", PM: "μ.μ."},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "π.μ.", PM: "μ.μ."},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "π.μ.", PM: "μ.μ."},
			},
		},
	}
}
