package prg

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, y 'mettas' d. MMMM", Long: "y 'mettas' d. MMMM", Medium: "dd.MM 'st'. y", Short: "dd.MM.yy"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "rag", Feb: "was", Mar: "pūl", Apr: "sak", May: "zal", Jun: "sīm", Jul: "līp", Aug: "dag", Sep: "sil", Oct: "spa", Nov: "lap", Dec: "sal"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "R", Feb: "W", Mar: "P", Apr: "S", May: "Z", Jun: "S", Jul: "L", Aug: "D", Sep: "S", Oct: "S", Nov: "L", Dec: "S"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "rags", Feb: "wassarins", Mar: "pūlis", Apr: "sakkis", May: "zallaws", Jun: "sīmenis", Jul: "līpa", Aug: "daggis", Sep: "sillins", Oct: "spallins", Nov: "lapkrūtis", Dec: "sallaws"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "nad", Mon: "pan", Tue: "wis", Wed: "pus", Thu: "ket", Fri: "pēn", Sat: "sab"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "N", Mon: "P", Tue: "W", Wed: "P", Thu: "K", Fri: "P", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "nadīli", Mon: "panadīli", Tue: "wisasīdis", Wed: "pussisawaiti", Thu: "ketwirtiks", Fri: "pēntniks", Sat: "sabattika"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "ankstāinan", PM: "pa pussideinan"},
			},
		},
	}
}
