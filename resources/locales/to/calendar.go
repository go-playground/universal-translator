package to

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/yy"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Sān", Feb: "Fēp", Mar: "Maʻa", Apr: "ʻEpe", May: "Mē", Jun: "Sun", Jul: "Siu", Aug: "ʻAok", Sep: "Sep", Oct: "ʻOka", Nov: "Nōv", Dec: "Tīs"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "S", Feb: "F", Mar: "M", Apr: "E", May: "M", Jun: "S", Jul: "S", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "T"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Sānuali", Feb: "Fēpueli", Mar: "Maʻasi", Apr: "ʻEpeleli", May: "Mē", Jun: "Sune", Jul: "Siulai", Aug: "ʻAokosi", Sep: "Sepitema", Oct: "ʻOkatopa", Nov: "Nōvema", Dec: "Tīsema"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Sāp", Mon: "Mōn", Tue: "Tūs", Wed: "Pul", Thu: "Tuʻa", Fri: "Fal", Sat: "Tok"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "P", Thu: "T", Fri: "F", Sat: "T"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "Sāp", Mon: "Mōn", Tue: "Tūs", Wed: "Pul", Thu: "Tuʻa", Fri: "Fal", Sat: "Tok"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Sāpate", Mon: "Mōnite", Tue: "Tūsite", Wed: "Pulelulu", Thu: "Tuʻapulelulu", Fri: "Falaite", Sat: "Tokonaki"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			},
		},
	}
}
