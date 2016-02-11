package mas

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Dal", Feb: "Ará", Mar: "Ɔɛn", Apr: "Doy", May: "Lép", Jun: "Rok", Jul: "Sás", Aug: "Bɔ́r", Sep: "Kús", Oct: "Gís", Nov: "Shʉ́", Dec: "Ntʉ́"},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Oladalʉ́", Feb: "Arát", Mar: "Ɔɛnɨ́ɔɨŋɔk", Apr: "Olodoyíóríê inkókúâ", May: "Oloilépūnyīē inkókúâ", Jun: "Kújúɔrɔk", Jul: "Mórusásin", Aug: "Ɔlɔ́ɨ́bɔ́rárɛ", Sep: "Kúshîn", Oct: "Olgísan", Nov: "Pʉshʉ́ka", Dec: "Ntʉ́ŋʉ́s"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Jpi", Mon: "Jtt", Tue: "Jnn", Wed: "Jtn", Thu: "Alh", Fri: "Iju", Sat: "Jmo"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "2", Mon: "3", Tue: "4", Wed: "5", Thu: "6", Fri: "7", Sat: "1"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Jumapílí", Mon: "Jumatátu", Tue: "Jumane", Wed: "Jumatánɔ", Thu: "Alaámisi", Fri: "Jumáa", Sat: "Jumamósi"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "Ɛnkakɛnyá", PM: "Ɛndámâ"},
			},
		},
	}
}
