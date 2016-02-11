package ff

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "sii", Feb: "col", Mar: "mbo", Apr: "see", May: "duu", Jun: "kor", Jul: "mor", Aug: "juk", Sep: "slt", Oct: "yar", Nov: "jol", Dec: "bow"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "s", Feb: "c", Mar: "m", Apr: "s", May: "d", Jun: "k", Jul: "m", Aug: "j", Sep: "s", Oct: "y", Nov: "j", Dec: "b"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "siilo", Feb: "colte", Mar: "mbooy", Apr: "seeɗto", May: "duujal", Jun: "korse", Jul: "morso", Aug: "juko", Sep: "siilto", Oct: "yarkomaa", Nov: "jolal", Dec: "bowte"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "dew", Mon: "aaɓ", Tue: "maw", Wed: "nje", Thu: "naa", Fri: "mwd", Sat: "hbi"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "d", Mon: "a", Tue: "m", Wed: "n", Thu: "n", Fri: "m", Sat: "h"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "dewo", Mon: "aaɓnde", Tue: "mawbaare", Wed: "njeslaare", Thu: "naasaande", Fri: "mawnde", Sat: "hoore-biir"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "subaka", PM: "kikiiɗe"},
			},
		},
	}
}
