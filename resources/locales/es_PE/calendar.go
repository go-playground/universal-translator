package es_PE

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: "d/MM/yy"},
			Time:     ut.CalendarDateFormat{},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Ene.", Feb: "Feb.", Mar: "Mar.", Apr: "Abr.", May: "May.", Jun: "Jun.", Jul: "Jul.", Aug: "Ago.", Sep: "Set.", Oct: "Oct.", Nov: "Nov.", Dec: "Dic."},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Enero", Feb: "Febrero", Mar: "Marzo", Apr: "Abril", May: "Mayo", Jun: "Junio", Jul: "Julio", Aug: "Agosto", Sep: "Setiembre", Oct: "Octubre", Nov: "Noviembre", Dec: "Diciembre"},
			},
			Days:    ut.CalendarDayFormatNames{},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
