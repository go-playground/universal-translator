package lt

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "y 'm'. MMMM d 'd'., EEEE", Long: "y 'm'. MMMM d 'd'.", Medium: "y-MM-dd", Short: "y-MM-dd"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "saus.", Feb: "vas.", Mar: "kov.", Apr: "bal.", May: "geg.", Jun: "birž.", Jul: "liep.", Aug: "rugp.", Sep: "rugs.", Oct: "spal.", Nov: "lapkr.", Dec: "gruod."},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "S", Feb: "V", Mar: "K", Apr: "B", May: "G", Jun: "B", Jul: "L", Aug: "R", Sep: "R", Oct: "S", Nov: "L", Dec: "G"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "sausis", Feb: "vasaris", Mar: "kovas", Apr: "balandis", May: "gegužė", Jun: "birželis", Jul: "liepa", Aug: "rugpjūtis", Sep: "rugsėjis", Oct: "spalis", Nov: "lapkritis", Dec: "gruodis"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "sk", Mon: "pr", Tue: "an", Wed: "tr", Thu: "kt", Fri: "pn", Sat: "št"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "S", Mon: "P", Tue: "A", Wed: "T", Thu: "K", Fri: "P", Sat: "Š"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "Sk", Mon: "Pr", Tue: "An", Wed: "Tr", Thu: "Kt", Fri: "Pn", Sat: "Št"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "sekmadienis", Mon: "pirmadienis", Tue: "antradienis", Wed: "trečiadienis", Thu: "ketvirtadienis", Fri: "penktadienis", Sat: "šeštadienis"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "priešpiet", PM: "popiet"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "pr. p.", PM: "pop."},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "priešpiet", PM: "popiet"},
			},
		},
	}
}
