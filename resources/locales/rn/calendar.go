package rn

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Mut.", Feb: "Gas.", Mar: "Wer.", Apr: "Mat.", May: "Gic.", Jun: "Kam.", Jul: "Nya.", Aug: "Kan.", Sep: "Nze.", Oct: "Ukw.", Nov: "Ugu.", Dec: "Uku."},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Nzero", Feb: "Ruhuhuma", Mar: "Ntwarante", Apr: "Ndamukiza", May: "Rusama", Jun: "Ruheshi", Jul: "Mukakaro", Aug: "Nyandagaro", Sep: "Nyakanga", Oct: "Gitugutu", Nov: "Munyonyo", Dec: "Kigarama"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "cu.", Mon: "mbe.", Tue: "kab.", Wed: "gtu.", Thu: "kan.", Fri: "gnu.", Sat: "gnd."},
				Narrow:      ut.CalendarDayFormatNameValue{},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Ku w’indwi", Mon: "Ku wa mbere", Tue: "Ku wa kabiri", Wed: "Ku wa gatatu", Thu: "Ku wa kane", Fri: "Ku wa gatanu", Sat: "Ku wa gatandatu"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "Z.MU.", PM: "Z.MW."},
			},
		},
	}
}
