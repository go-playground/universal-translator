package ee

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, MMMM d 'lia' y", Long: "MMMM d 'lia' y", Medium: "MMM d 'lia', y", Short: "M/d/yy"},
			Time:     ut.CalendarDateFormat{Full: "a 'ga' h:mm:ss zzzz", Long: "a 'ga' h:mm:ss z", Medium: "a 'ga' h:mm:ss", Short: "a 'ga' h:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{0} {1}", Long: "{0} {1}", Medium: "{0} {1}", Short: "{0} {1}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "dzv", Feb: "dzd", Mar: "ted", Apr: "afɔ", May: "dam", Jun: "mas", Jul: "sia", Aug: "dea", Sep: "any", Oct: "kel", Nov: "ade", Dec: "dzm"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "d", Feb: "d", Mar: "t", Apr: "a", May: "d", Jun: "m", Jul: "s", Aug: "d", Sep: "a", Oct: "k", Nov: "a", Dec: "d"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "dzove", Feb: "dzodze", Mar: "tedoxe", Apr: "afɔfĩe", May: "dama", Jun: "masa", Jul: "siamlɔm", Aug: "deasiamime", Sep: "anyɔnyɔ", Oct: "kele", Nov: "adeɛmekpɔxe", Dec: "dzome"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "kɔs", Mon: "dzo", Tue: "bla", Wed: "kuɖ", Thu: "yaw", Fri: "fiɖ", Sat: "mem"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "k", Mon: "d", Tue: "b", Wed: "k", Thu: "y", Fri: "f", Sat: "m"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "kɔs", Mon: "dzo", Tue: "bla", Wed: "kuɖ", Thu: "yaw", Fri: "fiɖ", Sat: "mem"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "kɔsiɖa", Mon: "dzoɖa", Tue: "blaɖa", Wed: "kuɖa", Thu: "yawoɖa", Fri: "fiɖa", Sat: "memleɖa"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "ŋ", PM: "ɣ"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "ŋdi", PM: "ɣetrɔ"},
			},
		},
	}
}
