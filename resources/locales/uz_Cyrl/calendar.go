package uz_Cyrl

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, y MMMM dd", Long: "y MMMM d", Medium: "y MMM d", Short: "dd/MM/yy"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Янв", Feb: "Фев", Mar: "Мар", Apr: "Апр", May: "Май", Jun: "Июн", Jul: "Июл", Aug: "Авг", Sep: "Сен", Oct: "Окт", Nov: "Ноя", Dec: "Дек"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "Я", Feb: "Ф", Mar: "М", Apr: "А", May: "М", Jun: "И", Jul: "И", Aug: "А", Sep: "С", Oct: "О", Nov: "Н", Dec: "Д"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Январ", Feb: "Феврал", Mar: "Март", Apr: "Апрел", May: "Май", Jun: "Июн", Jul: "Июл", Aug: "Август", Sep: "Сентябр", Oct: "Октябр", Nov: "Ноябр", Dec: "Декабр"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Якш", Mon: "Душ", Tue: "Сеш", Wed: "Чор", Thu: "Пай", Fri: "Жум", Sat: "Шан"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "Я", Mon: "Д", Tue: "С", Wed: "Ч", Thu: "П", Fri: "Ж", Sat: "Ш"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "Якш", Mon: "Душ", Tue: "Сеш", Wed: "Чор", Thu: "Пай", Fri: "Жум", Sat: "Шан"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "якшанба", Mon: "душанба", Tue: "сешанба", Wed: "чоршанба", Thu: "пайшанба", Fri: "жума", Sat: "шанба"},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
