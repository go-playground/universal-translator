package ak

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, y MMMM dd", Long: "y MMMM d", Medium: "y MMM d", Short: "yy/MM/dd"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "S-Ɔ", Feb: "K-Ɔ", Mar: "E-Ɔ", Apr: "E-O", May: "E-K", Jun: "O-A", Jul: "A-K", Aug: "D-Ɔ", Sep: "F-Ɛ", Oct: "Ɔ-A", Nov: "Ɔ-O", Dec: "M-Ɔ"},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Sanda-Ɔpɛpɔn", Feb: "Kwakwar-Ɔgyefuo", Mar: "Ebɔw-Ɔbenem", Apr: "Ebɔbira-Oforisuo", May: "Esusow Aketseaba-Kɔtɔnimba", Jun: "Obirade-Ayɛwohomumu", Jul: "Ayɛwoho-Kitawonsa", Aug: "Difuu-Ɔsandaa", Sep: "Fankwa-Ɛbɔ", Oct: "Ɔbɛsɛ-Ahinime", Nov: "Ɔberɛfɛw-Obubuo", Dec: "Mumu-Ɔpɛnimba"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Kwe", Mon: "Dwo", Tue: "Ben", Wed: "Wuk", Thu: "Yaw", Fri: "Fia", Sat: "Mem"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "K", Mon: "D", Tue: "B", Wed: "W", Thu: "Y", Fri: "F", Sat: "M"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Kwesida", Mon: "Dwowda", Tue: "Benada", Wed: "Wukuda", Thu: "Yawda", Fri: "Fida", Sat: "Memeneda"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "AN", PM: "EW"},
			},
		},
	}
}
