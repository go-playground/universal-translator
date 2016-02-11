package naq

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mar", Apr: "Apr", May: "May", Jun: "Jun", Jul: "Jul", Aug: "Aug", Sep: "Sep", Oct: "Oct", Nov: "Nov", Dec: "Dec"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "ǃKhanni", Feb: "ǃKhanǀgôab", Mar: "ǀKhuuǁkhâb", Apr: "ǃHôaǂkhaib", May: "ǃKhaitsâb", Jun: "Gamaǀaeb", Jul: "ǂKhoesaob", Aug: "Aoǁkhuumûǁkhâb", Sep: "Taraǀkhuumûǁkhâb", Oct: "ǂNûǁnâiseb", Nov: "ǀHooǂgaeb", Dec: "Hôasoreǁkhâb"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Son", Mon: "Ma", Tue: "De", Wed: "Wu", Thu: "Do", Fri: "Fr", Sat: "Sat"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "E", Wed: "W", Thu: "D", Fri: "F", Sat: "A"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Sontaxtsees", Mon: "Mantaxtsees", Tue: "Denstaxtsees", Wed: "Wunstaxtsees", Thu: "Dondertaxtsees", Fri: "Fraitaxtsees", Sat: "Satertaxtsees"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "ǁgoagas", PM: "ǃuias"},
			},
		},
	}
}
