package sq

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d.M.yy"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a, zzzz", Long: "h:mm:ss a, z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{Full: "{1} 'në' {0}", Long: "{1} 'në' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Shk", Mar: "Mar", Apr: "Pri", May: "Maj", Jun: "Qer", Jul: "Kor", Aug: "Gsh", Sep: "Sht", Oct: "Tet", Nov: "Nën", Dec: "Dhj"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "S", Mar: "M", Apr: "P", May: "M", Jun: "Q", Jul: "K", Aug: "G", Sep: "S", Oct: "T", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Janar", Feb: "Shkurt", Mar: "Mars", Apr: "Prill", May: "Maj", Jun: "Qershor", Jul: "Korrik", Aug: "Gusht", Sep: "Shtator", Oct: "Tetor", Nov: "Nëntor", Dec: "Dhjetor"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Die", Mon: "Hën", Tue: "Mar", Wed: "Mër", Thu: "Enj", Fri: "Pre", Sat: "Sht"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "D", Mon: "H", Tue: "M", Wed: "M", Thu: "E", Fri: "P", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "Die", Mon: "Hën", Tue: "Mar", Wed: "Mër", Thu: "Enj", Fri: "Pre", Sat: "Sht"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "E diel", Mon: "E hënë", Tue: "E martë", Wed: "E mërkurë", Thu: "E enjte", Fri: "E premte", Sat: "E shtunë"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "paradite", PM: "pasdite"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "paradite", PM: "pasdite"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "paradite", PM: "pasdite"},
			},
		},
	}
}
