package fi

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "cccc d. MMMM y", Long: "d. MMMM y", Medium: "d.M.y", Short: "d.M.y"},
			Time:     ut.CalendarDateFormat{Full: "H.mm.ss zzzz", Long: "H.mm.ss z", Medium: "H.mm.ss", Short: "H.mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} 'klo' {0}", Long: "{1} 'klo' {0}", Medium: "{1} 'klo' {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "tammi", Feb: "helmi", Mar: "maalis", Apr: "huhti", May: "touko", Jun: "kesä", Jul: "heinä", Aug: "elo", Sep: "syys", Oct: "loka", Nov: "marras", Dec: "joulu"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "T", Feb: "H", Mar: "M", Apr: "H", May: "T", Jun: "K", Jul: "H", Aug: "E", Sep: "S", Oct: "L", Nov: "M", Dec: "J"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "tammikuu", Feb: "helmikuu", Mar: "maaliskuu", Apr: "huhtikuu", May: "toukokuu", Jun: "kesäkuu", Jul: "heinäkuu", Aug: "elokuu", Sep: "syyskuu", Oct: "lokakuu", Nov: "marraskuu", Dec: "joulukuu"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "su", Mon: "ma", Tue: "ti", Wed: "ke", Thu: "to", Fri: "pe", Sat: "la"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "K", Thu: "T", Fri: "P", Sat: "L"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "su", Mon: "ma", Tue: "ti", Wed: "ke", Thu: "to", Fri: "pe", Sat: "la"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "sunnuntai", Mon: "maanantai", Tue: "tiistai", Wed: "keskiviikko", Thu: "torstai", Fri: "perjantai", Sat: "lauantai"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "ap.", PM: "ip."},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "ap.", PM: "ip."},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "ap.", PM: "ip."},
			},
		},
	}
}
