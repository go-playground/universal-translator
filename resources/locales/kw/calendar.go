package kw

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Gen", Feb: "Hwe", Mar: "Meu", Apr: "Ebr", May: "Me", Jun: "Met", Jul: "Gor", Aug: "Est", Sep: "Gwn", Oct: "Hed", Nov: "Du", Dec: "Kev"},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "mis Genver", Feb: "mis Hwevrer", Mar: "mis Meurth", Apr: "mis Ebrel", May: "mis Me", Jun: "mis Metheven", Jul: "mis Gortheren", Aug: "mis Est", Sep: "mis Gwynngala", Oct: "mis Hedra", Nov: "mis Du", Dec: "mis Kevardhu"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Sul", Mon: "Lun", Tue: "Mth", Wed: "Mhr", Thu: "Yow", Fri: "Gwe", Sat: "Sad"},
				Narrow:      ut.CalendarDayFormatNameValue{},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "dy Sul", Mon: "dy Lun", Tue: "dy Meurth", Wed: "dy Merher", Thu: "dy Yow", Fri: "dy Gwener", Sat: "dy Sadorn"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
			},
		},
	}
}
