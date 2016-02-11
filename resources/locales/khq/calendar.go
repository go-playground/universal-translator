package khq

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Žan", Feb: "Fee", Mar: "Mar", Apr: "Awi", May: "Me", Jun: "Žuw", Jul: "Žuy", Aug: "Ut", Sep: "Sek", Oct: "Okt", Nov: "Noo", Dec: "Dee"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "Ž", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "Ž", Jul: "Ž", Aug: "U", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Žanwiye", Feb: "Feewiriye", Mar: "Marsi", Apr: "Awiril", May: "Me", Jun: "Žuweŋ", Jul: "Žuyye", Aug: "Ut", Sep: "Sektanbur", Oct: "Oktoobur", Nov: "Noowanbur", Dec: "Deesanbur"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Alh", Mon: "Ati", Tue: "Ata", Wed: "Ala", Thu: "Alm", Fri: "Alj", Sat: "Ass"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "H", Mon: "T", Tue: "T", Wed: "L", Thu: "L", Fri: "L", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Alhadi", Mon: "Atini", Tue: "Atalata", Wed: "Alarba", Thu: "Alhamiisa", Fri: "Aljuma", Sat: "Assabdu"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "Adduha", PM: "Aluula"},
			},
		},
	}
}
