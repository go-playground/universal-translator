package tzm

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
			Time:     ut.CalendarDateFormat{},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Yen", Feb: "Yeb", Mar: "Mar", Apr: "Ibr", May: "May", Jun: "Yun", Jul: "Yul", Aug: "Ɣuc", Sep: "Cut", Oct: "Kṭu", Nov: "Nwa", Dec: "Duj"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "Y", Feb: "Y", Mar: "M", Apr: "I", May: "M", Jun: "Y", Jul: "Y", Aug: "Ɣ", Sep: "C", Oct: "K", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Yennayer", Feb: "Yebrayer", Mar: "Mars", Apr: "Ibrir", May: "Mayyu", Jun: "Yunyu", Jul: "Yulyuz", Aug: "Ɣuct", Sep: "Cutanbir", Oct: "Kṭuber", Nov: "Nwanbir", Dec: "Dujanbir"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Asa", Mon: "Ayn", Tue: "Asn", Wed: "Akr", Thu: "Akw", Fri: "Asm", Sat: "Asḍ"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "A", Mon: "A", Tue: "A", Wed: "A", Thu: "A", Fri: "A", Sat: "A"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Asamas", Mon: "Aynas", Tue: "Asinas", Wed: "Akras", Thu: "Akwas", Fri: "Asimwas", Sat: "Asiḍyas"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "Zdat azal", PM: "Ḍeffir aza"},
			},
		},
	}
}
