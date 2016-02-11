package pl

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "dd.MM.y", Short: "dd.MM.y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "sty", Feb: "lut", Mar: "mar", Apr: "kwi", May: "maj", Jun: "cze", Jul: "lip", Aug: "sie", Sep: "wrz", Oct: "paź", Nov: "lis", Dec: "gru"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "s", Feb: "l", Mar: "m", Apr: "k", May: "m", Jun: "c", Jul: "l", Aug: "s", Sep: "w", Oct: "p", Nov: "l", Dec: "g"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "styczeń", Feb: "luty", Mar: "marzec", Apr: "kwiecień", May: "maj", Jun: "czerwiec", Jul: "lipiec", Aug: "sierpień", Sep: "wrzesień", Oct: "październik", Nov: "listopad", Dec: "grudzień"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "niedz.", Mon: "pon.", Tue: "wt.", Wed: "śr.", Thu: "czw.", Fri: "pt.", Sat: "sob."},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "N", Mon: "P", Tue: "W", Wed: "Ś", Thu: "C", Fri: "P", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "niedz.", Mon: "pon.", Tue: "wt.", Wed: "śr.", Thu: "czw.", Fri: "pt.", Sat: "sob."},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "niedziela", Mon: "poniedziałek", Tue: "wtorek", Wed: "środa", Thu: "czwartek", Fri: "piątek", Sat: "sobota"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			},
		},
	}
}
