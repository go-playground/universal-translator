package xog

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mar", Apr: "Apu", May: "Maa", Jun: "Juu", Jul: "Jul", Aug: "Agu", Sep: "Seb", Oct: "Oki", Nov: "Nov", Dec: "Des"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Janwaliyo", Feb: "Febwaliyo", Mar: "Marisi", Apr: "Apuli", May: "Maayi", Jun: "Juuni", Jul: "Julaayi", Aug: "Agusito", Sep: "Sebuttemba", Oct: "Okitobba", Nov: "Novemba", Dec: "Desemba"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Sabi", Mon: "Bala", Tue: "Kubi", Wed: "Kusa", Thu: "Kuna", Fri: "Kuta", Sat: "Muka"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "S", Mon: "B", Tue: "B", Wed: "S", Thu: "K", Fri: "K", Sat: "M"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Sabiiti", Mon: "Balaza", Tue: "Owokubili", Wed: "Owokusatu", Thu: "Olokuna", Fri: "Olokutaanu", Sat: "Olomukaaga"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "Munkyo", PM: "Eigulo"},
			},
		},
	}
}
