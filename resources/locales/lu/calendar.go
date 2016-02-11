package lu

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Cio", Feb: "Lui", Mar: "Lus", Apr: "Muu", May: "Lum", Jun: "Luf", Jul: "Kab", Aug: "Lush", Sep: "Lut", Oct: "Lun", Nov: "Kas", Dec: "Cis"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "C", Feb: "L", Mar: "L", Apr: "M", May: "L", Jun: "L", Jul: "K", Aug: "L", Sep: "L", Oct: "L", Nov: "K", Dec: "C"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Ciongo", Feb: "Lùishi", Mar: "Lusòlo", Apr: "Mùuyà", May: "Lumùngùlù", Jun: "Lufuimi", Jul: "Kabàlàshìpù", Aug: "Lùshìkà", Sep: "Lutongolo", Oct: "Lungùdi", Nov: "Kaswèkèsè", Dec: "Ciswà"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Lum", Mon: "Nko", Tue: "Ndy", Wed: "Ndg", Thu: "Njw", Fri: "Ngv", Sat: "Lub"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "L", Mon: "N", Tue: "N", Wed: "N", Thu: "N", Fri: "N", Sat: "L"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Lumingu", Mon: "Nkodya", Tue: "Ndàayà", Wed: "Ndangù", Thu: "Njòwa", Fri: "Ngòvya", Sat: "Lubingu"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "Dinda", PM: "Dilolo"},
			},
		},
	}
}
