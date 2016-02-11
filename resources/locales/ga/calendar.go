package ga

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Ean", Feb: "Feabh", Mar: "Márta", Apr: "Aib", May: "Beal", Jun: "Meith", Jul: "Iúil", Aug: "Lún", Sep: "MFómh", Oct: "DFómh", Nov: "Samh", Dec: "Noll"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "E", Feb: "F", Mar: "M", Apr: "A", May: "B", Jun: "M", Jul: "I", Aug: "L", Sep: "M", Oct: "D", Nov: "S", Dec: "N"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Eanáir", Feb: "Feabhra", Mar: "Márta", Apr: "Aibreán", May: "Bealtaine", Jun: "Meitheamh", Jul: "Iúil", Aug: "Lúnasa", Sep: "Meán Fómhair", Oct: "Deireadh Fómhair", Nov: "Samhain", Dec: "Nollaig"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Domh", Mon: "Luan", Tue: "Máirt", Wed: "Céad", Thu: "Déar", Fri: "Aoine", Sat: "Sath"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "C", Thu: "D", Fri: "A", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "Do", Mon: "Lu", Tue: "Má", Wed: "Cé", Thu: "Dé", Fri: "Ao", Sat: "Sa"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Dé Domhnaigh", Mon: "Dé Luain", Tue: "Dé Máirt", Wed: "Dé Céadaoin", Thu: "Déardaoin", Fri: "Dé hAoine", Sat: "Dé Sathairn"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
			},
		},
	}
}
