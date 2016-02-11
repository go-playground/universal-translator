package seh

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d 'de' MMM 'de' y", Short: "d/M/y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Fev", Mar: "Mar", Apr: "Abr", May: "Mai", Jun: "Jun", Jul: "Jul", Aug: "Aug", Sep: "Set", Oct: "Otu", Nov: "Nov", Dec: "Dec"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Janeiro", Feb: "Fevreiro", Mar: "Marco", Apr: "Abril", May: "Maio", Jun: "Junho", Jul: "Julho", Aug: "Augusto", Sep: "Setembro", Oct: "Otubro", Nov: "Novembro", Dec: "Decembro"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Dim", Mon: "Pos", Tue: "Pir", Wed: "Tat", Thu: "Nai", Fri: "Sha", Sat: "Sab"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "D", Mon: "P", Tue: "C", Wed: "T", Thu: "N", Fri: "S", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Dimingu", Mon: "Chiposi", Tue: "Chipiri", Wed: "Chitatu", Thu: "Chinai", Fri: "Chishanu", Sat: "Sabudu"},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
