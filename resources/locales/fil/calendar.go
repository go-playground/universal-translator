package fil

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{Full: "{1} 'nang' {0}", Long: "{1} 'nang' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Ene", Feb: "Peb", Mar: "Mar", Apr: "Abr", May: "May", Jun: "Hun", Jul: "Hul", Aug: "Ago", Sep: "Set", Oct: "Okt", Nov: "Nob", Dec: "Dis"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "E", Feb: "P", Mar: "M", Apr: "A", May: "M", Jun: "Hun", Jul: "Hul", Aug: "Ago", Sep: "Set", Oct: "Okt", Nov: "Nob", Dec: "Dis"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Enero", Feb: "Pebrero", Mar: "Marso", Apr: "Abril", May: "Mayo", Jun: "Hunyo", Jul: "Hulyo", Aug: "Agosto", Sep: "Setyembre", Oct: "Oktubre", Nov: "Nobyembre", Dec: "Disyembre"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Lin", Mon: "Lun", Tue: "Mar", Wed: "Miy", Thu: "Huw", Fri: "Biy", Sat: "Sab"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "Lin", Mon: "Lun", Tue: "Mar", Wed: "Miy", Thu: "Huw", Fri: "Biy", Sat: "Sab"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "Li", Mon: "Lu", Tue: "Ma", Wed: "Mi", Thu: "Hu", Fri: "Bi", Sat: "Sa"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Linggo", Mon: "Lunes", Tue: "Martes", Wed: "Miyerkules", Thu: "Huwebes", Fri: "Biyernes", Sat: "Sabado"},
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
