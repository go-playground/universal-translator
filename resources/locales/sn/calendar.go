package sn

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Ndi", Feb: "Kuk", Mar: "Kur", Apr: "Kub", May: "Chv", Jun: "Chk", Jul: "Chg", Aug: "Nya", Sep: "Gun", Oct: "Gum", Nov: "Mb", Dec: "Zvi"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "N", Feb: "K", Mar: "K", Apr: "K", May: "C", Jun: "C", Jul: "C", Aug: "N", Sep: "G", Oct: "G", Nov: "M", Dec: "Z"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Ndira", Feb: "Kukadzi", Mar: "Kurume", Apr: "Kubvumbi", May: "Chivabvu", Jun: "Chikumi", Jul: "Chikunguru", Aug: "Nyamavhuvhu", Sep: "Gunyana", Oct: "Gumiguru", Nov: "Mbudzi", Dec: "Zvita"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Svo", Mon: "Muv", Tue: "Chip", Wed: "Chit", Thu: "Chin", Fri: "Chis", Sat: "Mug"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "C", Wed: "C", Thu: "C", Fri: "C", Sat: "M"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Svondo", Mon: "Muvhuro", Tue: "Chipiri", Wed: "Chitatu", Thu: "China", Fri: "Chishanu", Sat: "Mugovera"},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
