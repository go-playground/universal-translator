package teo

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Rar", Feb: "Muk", Mar: "Kwa", Apr: "Dun", May: "Mar", Jun: "Mod", Jul: "Jol", Aug: "Ped", Sep: "Sok", Oct: "Tib", Nov: "Lab", Dec: "Poo"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "R", Feb: "M", Mar: "K", Apr: "D", May: "M", Jun: "M", Jul: "J", Aug: "P", Sep: "S", Oct: "T", Nov: "L", Dec: "P"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Orara", Feb: "Omuk", Mar: "Okwamg’", Apr: "Odung’el", May: "Omaruk", Jun: "Omodok’king’ol", Jul: "Ojola", Aug: "Opedel", Sep: "Osokosokoma", Oct: "Otibar", Nov: "Olabor", Dec: "Opoo"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Jum", Mon: "Bar", Tue: "Aar", Wed: "Uni", Thu: "Ung", Fri: "Kan", Sat: "Sab"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "J", Mon: "B", Tue: "A", Wed: "U", Thu: "U", Fri: "K", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Nakaejuma", Mon: "Nakaebarasa", Tue: "Nakaare", Wed: "Nakauni", Thu: "Nakaung’on", Fri: "Nakakany", Sat: "Nakasabiti"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "Taparachu", PM: "Ebongi"},
			},
		},
	}
}
