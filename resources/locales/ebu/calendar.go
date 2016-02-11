package ebu

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Mbe", Feb: "Kai", Mar: "Kat", Apr: "Kan", May: "Gat", Jun: "Gan", Jul: "Mug", Aug: "Knn", Sep: "Ken", Oct: "Iku", Nov: "Imw", Dec: "Igi"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "M", Feb: "K", Mar: "K", Apr: "K", May: "G", Jun: "G", Jul: "M", Aug: "K", Sep: "K", Oct: "I", Nov: "I", Dec: "I"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Mweri wa mbere", Feb: "Mweri wa kaĩri", Mar: "Mweri wa kathatũ", Apr: "Mweri wa kana", May: "Mweri wa gatano", Jun: "Mweri wa gatantatũ", Jul: "Mweri wa mũgwanja", Aug: "Mweri wa kanana", Sep: "Mweri wa kenda", Oct: "Mweri wa ikũmi", Nov: "Mweri wa ikũmi na ũmwe", Dec: "Mweri wa ikũmi na Kaĩrĩ"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Kma", Mon: "Tat", Tue: "Ine", Wed: "Tan", Thu: "Arm", Fri: "Maa", Sat: "NMM"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "K", Mon: "N", Tue: "N", Wed: "N", Thu: "A", Fri: "M", Sat: "N"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Kiumia", Mon: "Njumatatu", Tue: "Njumaine", Wed: "Njumatano", Thu: "Aramithi", Fri: "Njumaa", Sat: "NJumamothii"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "KI", PM: "UT"},
			},
		},
	}
}
