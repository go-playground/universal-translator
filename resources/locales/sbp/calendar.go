package sbp

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Mup", Feb: "Mwi", Mar: "Msh", Apr: "Mun", May: "Mag", Jun: "Muj", Jul: "Msp", Aug: "Mpg", Sep: "Mye", Oct: "Mok", Nov: "Mus", Dec: "Muh"},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Mupalangulwa", Feb: "Mwitope", Mar: "Mushende", Apr: "Munyi", May: "Mushende Magali", Jun: "Mujimbi", Jul: "Mushipepo", Aug: "Mupuguto", Sep: "Munyense", Oct: "Mokhu", Nov: "Musongandembwe", Dec: "Muhaano"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Mul", Mon: "Jtt", Tue: "Jnn", Wed: "Jtn", Thu: "Alh", Fri: "Iju", Sat: "Jmo"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "M", Mon: "J", Tue: "J", Wed: "J", Thu: "A", Fri: "I", Sat: "J"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Mulungu", Mon: "Jumatatu", Tue: "Jumanne", Wed: "Jumatano", Thu: "Alahamisi", Fri: "Ijumaa", Sat: "Jumamosi"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "Lwamilawu", PM: "Pashamihe"},
			},
		},
	}
}
