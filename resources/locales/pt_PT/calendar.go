package pt_PT

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "", Long: "", Medium: "dd/MM/y", Short: ""},
			Time:     ut.CalendarDateFormat{},
			DateTime: ut.CalendarDateFormat{Full: "{1} 'às' {0}", Long: "{1} 'às' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "domingo", Mon: "segunda", Tue: "terça", Wed: "quarta", Thu: "quinta", Fri: "sexta", Sat: "sábado"},
				Narrow:      ut.CalendarDayFormatNameValue{},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "manhã", PM: "tarde"},
			},
		},
	}
}
