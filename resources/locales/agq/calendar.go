package agq

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "nùm", Feb: "kɨz", Mar: "tɨd", Apr: "taa", May: "see", Jun: "nzu", Jul: "dum", Aug: "fɔe", Sep: "dzu", Oct: "lɔm", Nov: "kaa", Dec: "fwo"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "n", Feb: "k", Mar: "t", Apr: "t", May: "s", Jun: "z", Jul: "k", Aug: "f", Sep: "d", Oct: "l", Nov: "c", Dec: "f"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "ndzɔ̀ŋɔ̀nùm", Feb: "ndzɔ̀ŋɔ̀kƗ̀zùʔ", Mar: "ndzɔ̀ŋɔ̀tƗ̀dʉ̀ghà", Apr: "ndzɔ̀ŋɔ̀tǎafʉ̄ghā", May: "ndzɔ̀ŋèsèe", Jun: "ndzɔ̀ŋɔ̀nzùghò", Jul: "ndzɔ̀ŋɔ̀dùmlo", Aug: "ndzɔ̀ŋɔ̀kwîfɔ̀e", Sep: "ndzɔ̀ŋɔ̀tƗ̀fʉ̀ghàdzughù", Oct: "ndzɔ̀ŋɔ̀ghǔuwelɔ̀m", Nov: "ndzɔ̀ŋɔ̀chwaʔàkaa wo", Dec: "ndzɔ̀ŋèfwòo"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "nts", Mon: "kpa", Tue: "ghɔ", Wed: "tɔm", Thu: "ume", Fri: "ghɨ", Sat: "dzk"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "n", Mon: "k", Tue: "g", Wed: "t", Thu: "u", Fri: "g", Sat: "d"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "tsuʔntsɨ", Mon: "tsuʔukpà", Tue: "tsuʔughɔe", Wed: "tsuʔutɔ̀mlò", Thu: "tsuʔumè", Fri: "tsuʔughɨ̂m", Sat: "tsuʔndzɨkɔʔɔ"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "a.g", PM: "a.k"},
			},
		},
	}
}
