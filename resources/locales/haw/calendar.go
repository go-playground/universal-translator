package haw

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/yy"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Ian.", Feb: "Pep.", Mar: "Mal.", Apr: "ʻAp.", May: "Mei", Jun: "Iun.", Jul: "Iul.", Aug: "ʻAu.", Sep: "Kep.", Oct: "ʻOk.", Nov: "Now.", Dec: "Kek."},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Ianuali", Feb: "Pepeluali", Mar: "Malaki", Apr: "ʻApelila", May: "Mei", Jun: "Iune", Jul: "Iulai", Aug: "ʻAukake", Sep: "Kepakemapa", Oct: "ʻOkakopa", Nov: "Nowemapa", Dec: "Kekemapa"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "LP", Mon: "P1", Tue: "P2", Wed: "P3", Thu: "P4", Fri: "P5", Sat: "P6"},
				Narrow:      ut.CalendarDayFormatNameValue{},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Lāpule", Mon: "Poʻakahi", Tue: "Poʻalua", Wed: "Poʻakolu", Thu: "Poʻahā", Fri: "Poʻalima", Sat: "Poʻaono"},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
