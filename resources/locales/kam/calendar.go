package kam

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Mbe", Feb: "Kel", Mar: "Ktũ", Apr: "Kan", May: "Ktn", Jun: "Tha", Jul: "Moo", Aug: "Nya", Sep: "Knd", Oct: "Ĩku", Nov: "Ĩkm", Dec: "Ĩkl"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "M", Feb: "K", Mar: "K", Apr: "K", May: "K", Jun: "T", Jul: "M", Aug: "N", Sep: "K", Oct: "Ĩ", Nov: "Ĩ", Dec: "Ĩ"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Mwai wa mbee", Feb: "Mwai wa kelĩ", Mar: "Mwai wa katatũ", Apr: "Mwai wa kana", May: "Mwai wa katano", Jun: "Mwai wa thanthatũ", Jul: "Mwai wa muonza", Aug: "Mwai wa nyaanya", Sep: "Mwai wa kenda", Oct: "Mwai wa ĩkumi", Nov: "Mwai wa ĩkumi na ĩmwe", Dec: "Mwai wa ĩkumi na ilĩ"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Wky", Mon: "Wkw", Tue: "Wkl", Wed: "Wtũ", Thu: "Wkn", Fri: "Wtn", Sat: "Wth"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "Y", Mon: "W", Tue: "E", Wed: "A", Thu: "A", Fri: "A", Sat: "A"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Wa kyumwa", Mon: "Wa kwambĩlĩlya", Tue: "Wa kelĩ", Wed: "Wa katatũ", Thu: "Wa kana", Fri: "Wa katano", Sat: "Wa thanthatũ"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "Ĩyakwakya", PM: "Ĩyawĩoo"},
			},
		},
	}
}
