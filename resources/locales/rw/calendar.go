package rw

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, y MMMM dd", Long: "y MMMM d", Medium: "y MMM d", Short: "yy/MM/dd"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "mut.", Feb: "gas.", Mar: "wer.", Apr: "mat.", May: "gic.", Jun: "kam.", Jul: "nya.", Aug: "kan.", Sep: "nze.", Oct: "ukw.", Nov: "ugu.", Dec: "uku."},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Mutarama", Feb: "Gashyantare", Mar: "Werurwe", Apr: "Mata", May: "Gicuransi", Jun: "Kamena", Jul: "Nyakanga", Aug: "Kanama", Sep: "Nzeli", Oct: "Ukwakira", Nov: "Ugushyingo", Dec: "Ukuboza"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "cyu.", Mon: "mbe.", Tue: "kab.", Wed: "gtu.", Thu: "kan.", Fri: "gnu.", Sat: "gnd."},
				Narrow:      ut.CalendarDayFormatNameValue{},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Ku cyumweru", Mon: "Kuwa mbere", Tue: "Kuwa kabiri", Wed: "Kuwa gatatu", Thu: "Kuwa kane", Fri: "Kuwa gatanu", Sat: "Kuwa gatandatu"},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
