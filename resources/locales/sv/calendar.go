package sv

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "y-MM-dd"},
			Time:     ut.CalendarDateFormat{Full: "'kl'. HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "jan.", Feb: "feb.", Mar: "mars", Apr: "apr.", May: "maj", Jun: "juni", Jul: "juli", Aug: "aug.", Sep: "sep.", Oct: "okt.", Nov: "nov.", Dec: "dec."},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "januari", Feb: "februari", Mar: "mars", Apr: "april", May: "maj", Jun: "juni", Jul: "juli", Aug: "augusti", Sep: "september", Oct: "oktober", Nov: "november", Dec: "december"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "sön", Mon: "mån", Tue: "tis", Wed: "ons", Thu: "tors", Fri: "fre", Sat: "lör"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "O", Thu: "T", Fri: "F", Sat: "L"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "sö", Mon: "må", Tue: "ti", Wed: "on", Thu: "to", Fri: "fr", Sat: "lö"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "söndag", Mon: "måndag", Tue: "tisdag", Wed: "onsdag", Thu: "torsdag", Fri: "fredag", Sat: "lördag"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "f.m.", PM: "e.m."},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "fm", PM: "em"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "förmiddag", PM: "eftermiddag"},
			},
		},
	}
}
