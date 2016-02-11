package es

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d MMM y", Short: "d/M/yy"},
			Time:     ut.CalendarDateFormat{Full: "H:mm:ss (zzzz)", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "ene.", Feb: "feb.", Mar: "mar.", Apr: "abr.", May: "may.", Jun: "jun.", Jul: "jul.", Aug: "ago.", Sep: "sept.", Oct: "oct.", Nov: "nov.", Dec: "dic."},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "E", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "enero", Feb: "febrero", Mar: "marzo", Apr: "abril", May: "mayo", Jun: "junio", Jul: "julio", Aug: "agosto", Sep: "septiembre", Oct: "octubre", Nov: "noviembre", Dec: "diciembre"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "dom.", Mon: "lun.", Tue: "mar.", Wed: "mié.", Thu: "jue.", Fri: "vie.", Sat: "sáb."},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "X", Thu: "J", Fri: "V", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "DO", Mon: "LU", Tue: "MA", Wed: "MI", Thu: "JU", Fri: "VI", Sat: "SA"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "domingo", Mon: "lunes", Tue: "martes", Wed: "miércoles", Thu: "jueves", Fri: "viernes", Sat: "sábado"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "a. m.", PM: "p. m."},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "a. m.", PM: "p. m."},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "a. m.", PM: "p. m."},
			},
		},
	}
}
