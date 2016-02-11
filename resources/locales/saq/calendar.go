package saq

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Obo", Feb: "Waa", Mar: "Oku", Apr: "Ong", May: "Ime", Jun: "Ile", Jul: "Sap", Aug: "Isi", Sep: "Saa", Oct: "Tom", Nov: "Tob", Dec: "Tow"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "O", Feb: "W", Mar: "O", Apr: "O", May: "I", Jun: "I", Jul: "S", Aug: "I", Sep: "S", Oct: "T", Nov: "T", Dec: "T"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Lapa le obo", Feb: "Lapa le waare", Mar: "Lapa le okuni", Apr: "Lapa le ong’wan", May: "Lapa le imet", Jun: "Lapa le ile", Jul: "Lapa le sapa", Aug: "Lapa le isiet", Sep: "Lapa le saal", Oct: "Lapa le tomon", Nov: "Lapa le tomon obo", Dec: "Lapa le tomon waare"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Are", Mon: "Kun", Tue: "Ong", Wed: "Ine", Thu: "Ile", Fri: "Sap", Sat: "Kwe"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "A", Mon: "K", Tue: "O", Wed: "I", Thu: "I", Fri: "S", Sat: "K"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Mderot ee are", Mon: "Mderot ee kuni", Tue: "Mderot ee ong’wan", Wed: "Mderot ee inet", Thu: "Mderot ee ile", Fri: "Mderot ee sapa", Sat: "Mderot ee kwe"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "Tesiran", PM: "Teipa"},
			},
		},
	}
}
