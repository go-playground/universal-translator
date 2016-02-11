package se

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "ođđj", Feb: "guov", Mar: "njuk", Apr: "cuo", May: "mies", Jun: "geas", Jul: "suoi", Aug: "borg", Sep: "čakč", Oct: "golg", Nov: "skáb", Dec: "juov"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "O", Feb: "G", Mar: "N", Apr: "C", May: "M", Jun: "G", Jul: "S", Aug: "B", Sep: "Č", Oct: "G", Nov: "S", Dec: "J"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "ođđajagemánnu", Feb: "guovvamánnu", Mar: "njukčamánnu", Apr: "cuoŋománnu", May: "miessemánnu", Jun: "geassemánnu", Jul: "suoidnemánnu", Aug: "borgemánnu", Sep: "čakčamánnu", Oct: "golggotmánnu", Nov: "skábmamánnu", Dec: "juovlamánnu"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "sotn", Mon: "vuos", Tue: "maŋ", Wed: "gask", Thu: "duor", Fri: "bear", Sat: "láv"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "S", Mon: "V", Tue: "M", Wed: "G", Thu: "D", Fri: "B", Sat: "L"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "sotnabeaivi", Mon: "vuossárga", Tue: "maŋŋebárga", Wed: "gaskavahkku", Thu: "duorasdat", Fri: "bearjadat", Sat: "lávvardat"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "i.b.", PM: "e.b."},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "iđitbeaivi", PM: "eahketbeaivi"},
			},
		},
	}
}
