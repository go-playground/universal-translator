package fur

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE d 'di' MMMM 'dal' y", Long: "d 'di' MMMM 'dal' y", Medium: "dd/MM/y", Short: "dd/MM/yy"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Zen", Feb: "Fev", Mar: "Mar", Apr: "Avr", May: "Mai", Jun: "Jug", Jul: "Lui", Aug: "Avo", Sep: "Set", Oct: "Otu", Nov: "Nov", Dec: "Dic"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "Z", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "L", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Zenâr", Feb: "Fevrâr", Mar: "Març", Apr: "Avrîl", May: "Mai", Jun: "Jugn", Jul: "Lui", Aug: "Avost", Sep: "Setembar", Oct: "Otubar", Nov: "Novembar", Dec: "Dicembar"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "dom", Mon: "lun", Tue: "mar", Wed: "mie", Thu: "joi", Fri: "vin", Sat: "sab"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "M", Thu: "J", Fri: "V", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "domenie", Mon: "lunis", Tue: "martars", Wed: "miercus", Thu: "joibe", Fri: "vinars", Sat: "sabide"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "a.", PM: "p."},
			},
		},
	}
}
