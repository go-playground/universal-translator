package be

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d.M.y", Short: "d.M.yy"},
			Time:     ut.CalendarDateFormat{Full: "HH.mm.ss zzzz", Long: "HH.mm.ss z", Medium: "HH.mm.ss", Short: "HH.mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} 'у' {0}", Long: "{1} 'у' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "сту", Feb: "лют", Mar: "сак", Apr: "кра", May: "май", Jun: "чэр", Jul: "ліп", Aug: "жні", Sep: "вер", Oct: "кас", Nov: "ліс", Dec: "сне"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "с", Feb: "л", Mar: "с", Apr: "к", May: "м", Jun: "ч", Jul: "л", Aug: "ж", Sep: "в", Oct: "к", Nov: "л", Dec: "с"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "студзень", Feb: "люты", Mar: "сакавік", Apr: "красавік", May: "май", Jun: "чэрвень", Jul: "ліпень", Aug: "жнівень", Sep: "верасень", Oct: "кастрычнік", Nov: "лістапад", Dec: "снежань"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "нд", Mon: "пн", Tue: "аў", Wed: "ср", Thu: "чц", Fri: "пт", Sat: "сб"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "н", Mon: "п", Tue: "а", Wed: "с", Thu: "ч", Fri: "п", Sat: "с"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "нд", Mon: "пн", Tue: "аў", Wed: "ср", Thu: "чц", Fri: "пт", Sat: "сб"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "нядзеля", Mon: "панядзелак", Tue: "аўторак", Wed: "серада", Thu: "чацвер", Fri: "пятніца", Sat: "субота"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "раніцы", PM: "вечара"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "раніцы", PM: "вечара"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "да паўдня", PM: "пасля паўдня"},
			},
		},
	}
}
