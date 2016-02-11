package ca

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM 'de' y", Long: "d MMMM 'de' y", Medium: "d MMM y", Short: "d/M/yy"},
			Time:     ut.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} 'a' 'les' {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "gen.", Feb: "febr.", Mar: "març", Apr: "abr.", May: "maig", Jun: "juny", Jul: "jul.", Aug: "ag.", Sep: "set.", Oct: "oct.", Nov: "nov.", Dec: "des."},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "GN", Feb: "FB", Mar: "MÇ", Apr: "AB", May: "MG", Jun: "JN", Jul: "JL", Aug: "AG", Sep: "ST", Oct: "OC", Nov: "NV", Dec: "DS"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "gener", Feb: "febrer", Mar: "març", Apr: "abril", May: "maig", Jun: "juny", Jul: "juliol", Aug: "agost", Sep: "setembre", Oct: "octubre", Nov: "novembre", Dec: "desembre"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "dg.", Mon: "dl.", Tue: "dt.", Wed: "dc.", Thu: "dj.", Fri: "dv.", Sat: "ds."},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "dg", Mon: "dl", Tue: "dt", Wed: "dc", Thu: "dj", Fri: "dv", Sat: "ds"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "dg.", Mon: "dl.", Tue: "dt.", Wed: "dc.", Thu: "dj.", Fri: "dv.", Sat: "ds."},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "diumenge", Mon: "dilluns", Tue: "dimarts", Wed: "dimecres", Thu: "dijous", Fri: "divendres", Sat: "dissabte"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "a. m.", PM: "p. m."},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "a. m.", PM: "p. m."},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "a. m.", PM: "p. m."},
			},
		},
	}
}
