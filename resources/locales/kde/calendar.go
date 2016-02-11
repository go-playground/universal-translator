package kde

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mac", Apr: "Apr", May: "Mei", Jun: "Jun", Jul: "Jul", Aug: "Ago", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Des"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Mwedi Ntandi", Feb: "Mwedi wa Pili", Mar: "Mwedi wa Tatu", Apr: "Mwedi wa Nchechi", May: "Mwedi wa Nnyano", Jun: "Mwedi wa Nnyano na Umo", Jul: "Mwedi wa Nnyano na Mivili", Aug: "Mwedi wa Nnyano na Mitatu", Sep: "Mwedi wa Nnyano na Nchechi", Oct: "Mwedi wa Nnyano na Nnyano", Nov: "Mwedi wa Nnyano na Nnyano na U", Dec: "Mwedi wa Nnyano na Nnyano na M"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Ll2", Mon: "Ll3", Tue: "Ll4", Wed: "Ll5", Thu: "Ll6", Fri: "Ll7", Sat: "Ll1"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "2", Mon: "3", Tue: "4", Wed: "5", Thu: "6", Fri: "7", Sat: "1"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Liduva lyapili", Mon: "Liduva lyatatu", Tue: "Liduva lyanchechi", Wed: "Liduva lyannyano", Thu: "Liduva lyannyano na linji", Fri: "Liduva lyannyano na mavili", Sat: "Liduva litandi"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "Muhi", PM: "Chilo"},
			},
		},
	}
}
