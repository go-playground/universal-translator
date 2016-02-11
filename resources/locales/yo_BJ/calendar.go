package yo_BJ

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Shɛ́rɛ́", Feb: "Èrèlè", Mar: "Ɛrɛ̀nà", Apr: "Ìgbé", May: "Ɛ̀bibi", Jun: "Òkúdu", Jul: "Agɛmɔ", Aug: "Ògún", Sep: "Owewe", Oct: "Ɔ̀wàrà", Nov: "Bélú", Dec: "Ɔ̀pɛ̀"},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Oshù Shɛ́rɛ́", Feb: "Oshù Èrèlè", Mar: "Oshù Ɛrɛ̀nà", Apr: "Oshù Ìgbé", May: "Oshù Ɛ̀bibi", Jun: "Oshù Òkúdu", Jul: "Oshù Agɛmɔ", Aug: "Oshù Ògún", Sep: "Oshù Owewe", Oct: "Oshù Ɔ̀wàrà", Nov: "Oshù Bélú", Dec: "Oshù Ɔ̀pɛ̀"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Àìkú", Mon: "Ajé", Tue: "Ìsɛ́gun", Wed: "Ɔjɔ́rú", Thu: "Ɔjɔ́bɔ", Fri: "Ɛtì", Sat: "Àbámɛ́ta"},
				Narrow:      ut.CalendarDayFormatNameValue{},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Ɔjɔ́ Àìkú", Mon: "Ɔjɔ́ Ajé", Tue: "Ɔjɔ́ Ìsɛ́gun", Wed: "Ɔjɔ́rú", Thu: "Ɔjɔ́bɔ", Fri: "Ɔjɔ́ Ɛtì", Sat: "Ɔjɔ́ Àbámɛ́ta"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "Àárɔ̀", PM: "Ɔ̀sán"},
			},
		},
	}
}
