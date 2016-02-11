package tr

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "d MMMM y EEEE", Long: "d MMMM y", Medium: "d MMM y", Short: "d.MM.y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Oca", Feb: "Şub", Mar: "Mar", Apr: "Nis", May: "May", Jun: "Haz", Jul: "Tem", Aug: "Ağu", Sep: "Eyl", Oct: "Eki", Nov: "Kas", Dec: "Ara"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "O", Feb: "Ş", Mar: "M", Apr: "N", May: "M", Jun: "H", Jul: "T", Aug: "A", Sep: "E", Oct: "E", Nov: "K", Dec: "A"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Ocak", Feb: "Şubat", Mar: "Mart", Apr: "Nisan", May: "Mayıs", Jun: "Haziran", Jul: "Temmuz", Aug: "Ağustos", Sep: "Eylül", Oct: "Ekim", Nov: "Kasım", Dec: "Aralık"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Paz", Mon: "Pzt", Tue: "Sal", Wed: "Çar", Thu: "Per", Fri: "Cum", Sat: "Cmt"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "P", Mon: "P", Tue: "S", Wed: "Ç", Thu: "P", Fri: "C", Sat: "C"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "Pa", Mon: "Pt", Tue: "Sa", Wed: "Ça", Thu: "Pe", Fri: "Cu", Sat: "Ct"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Pazar", Mon: "Pazartesi", Tue: "Salı", Wed: "Çarşamba", Thu: "Perşembe", Fri: "Cuma", Sat: "Cumartesi"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "ÖÖ", PM: "ÖS"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "ÖÖ", PM: "ÖS"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "ÖÖ", PM: "ÖS"},
			},
		},
	}
}
