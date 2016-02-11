package lag

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Fúngatɨ", Feb: "Naanɨ", Mar: "Keenda", Apr: "Ikúmi", May: "Inyambala", Jun: "Idwaata", Jul: "Mʉʉnchɨ", Aug: "Vɨɨrɨ", Sep: "Saatʉ", Oct: "Inyi", Nov: "Saano", Dec: "Sasatʉ"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "F", Feb: "N", Mar: "K", Apr: "I", May: "I", Jun: "I", Jul: "M", Aug: "V", Sep: "S", Oct: "I", Nov: "S", Dec: "S"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Kʉfúngatɨ", Feb: "Kʉnaanɨ", Mar: "Kʉkeenda", Apr: "Kwiikumi", May: "Kwiinyambála", Jun: "Kwiidwaata", Jul: "Kʉmʉʉnchɨ", Aug: "Kʉvɨɨrɨ", Sep: "Kʉsaatʉ", Oct: "Kwiinyi", Nov: "Kʉsaano", Dec: "Kʉsasatʉ"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Píili", Mon: "Táatu", Tue: "Íne", Wed: "Táano", Thu: "Alh", Fri: "Ijm", Sat: "Móosi"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "P", Mon: "T", Tue: "E", Wed: "O", Thu: "A", Fri: "I", Sat: "M"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Jumapíiri", Mon: "Jumatátu", Tue: "Jumaíne", Wed: "Jumatáano", Thu: "Alamíisi", Fri: "Ijumáa", Sat: "Jumamóosi"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "TOO", PM: "MUU"},
			},
		},
	}
}
