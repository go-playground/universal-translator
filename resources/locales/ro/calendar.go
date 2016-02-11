package ro

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd.MM.y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "ian.", Feb: "feb.", Mar: "mar.", Apr: "apr.", May: "mai", Jun: "iun.", Jul: "iul.", Aug: "aug.", Sep: "sept.", Oct: "oct.", Nov: "nov.", Dec: "dec."},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "I", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "I", Jul: "I", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "ianuarie", Feb: "februarie", Mar: "martie", Apr: "aprilie", May: "mai", Jun: "iunie", Jul: "iulie", Aug: "august", Sep: "septembrie", Oct: "octombrie", Nov: "noiembrie", Dec: "decembrie"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "dum.", Mon: "lun.", Tue: "mar.", Wed: "mie.", Thu: "joi", Fri: "vin.", Sat: "sâm."},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "M", Thu: "J", Fri: "V", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "du", Mon: "lu", Tue: "ma", Wed: "mi", Thu: "jo", Fri: "vi", Sat: "sâ"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "duminică", Mon: "luni", Tue: "marți", Wed: "miercuri", Thu: "joi", Fri: "vineri", Sat: "sâmbătă"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."},
			},
		},
	}
}
