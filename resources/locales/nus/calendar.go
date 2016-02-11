package nus

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/MM/y"},
			Time:     ut.CalendarDateFormat{Full: "zzzz h:mm:ss a", Long: "z h:mm:ss a", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Tiop", Feb: "Pɛt", Mar: "Duɔ̱ɔ̱", Apr: "Guak", May: "Duä", Jun: "Kor", Jul: "Pay", Aug: "Thoo", Sep: "Tɛɛ", Oct: "Laa", Nov: "Kur", Dec: "Tid"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "T", Feb: "P", Mar: "D", Apr: "G", May: "D", Jun: "K", Jul: "P", Aug: "T", Sep: "T", Oct: "L", Nov: "K", Dec: "T"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Tiop thar pɛt", Feb: "Pɛt", Mar: "Duɔ̱ɔ̱ŋ", Apr: "Guak", May: "Duät", Jun: "Kornyoot", Jul: "Pay yie̱tni", Aug: "Tho̱o̱r", Sep: "Tɛɛr", Oct: "Laath", Nov: "Kur", Dec: "Tio̱p in di̱i̱t"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Cäŋ", Mon: "Jiec", Tue: "Rɛw", Wed: "Diɔ̱k", Thu: "Ŋuaan", Fri: "Dhieec", Sat: "Bäkɛl"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "C", Mon: "J", Tue: "R", Wed: "D", Thu: "Ŋ", Fri: "D", Sat: "B"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Cäŋ kuɔth", Mon: "Jiec la̱t", Tue: "Rɛw lätni", Wed: "Diɔ̱k lätni", Thu: "Ŋuaan lätni", Fri: "Dhieec lätni", Sat: "Bäkɛl lätni"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "RW", PM: "TŊ"},
			},
		},
	}
}
