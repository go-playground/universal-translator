package uk

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y 'р'.", Long: "d MMMM y 'р'.", Medium: "d MMM y 'р'.", Short: "dd.MM.yy"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} 'о' {0}", Long: "{1} 'о' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Січ", Feb: "Лют", Mar: "Бер", Apr: "Кві", May: "Тра", Jun: "Чер", Jul: "Лип", Aug: "Сер", Sep: "Вер", Oct: "Жов", Nov: "Лис", Dec: "Гру"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "С", Feb: "Л", Mar: "Б", Apr: "К", May: "Т", Jun: "Ч", Jul: "Л", Aug: "С", Sep: "В", Oct: "Ж", Nov: "Л", Dec: "Г"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "січень", Feb: "лютий", Mar: "березень", Apr: "квітень", May: "травень", Jun: "червень", Jul: "липень", Aug: "серпень", Sep: "вересень", Oct: "жовтень", Nov: "листопад", Dec: "грудень"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Нд", Mon: "Пн", Tue: "Вт", Wed: "Ср", Thu: "Чт", Fri: "Пт", Sat: "Сб"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "Н", Mon: "П", Tue: "В", Wed: "С", Thu: "Ч", Fri: "П", Sat: "С"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "Нд", Mon: "Пн", Tue: "Вт", Wed: "Ср", Thu: "Чт", Fri: "Пт", Sat: "Сб"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "неділя", Mon: "понеділок", Tue: "вівторок", Wed: "середа", Thu: "четвер", Fri: "пʼятниця", Sat: "субота"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "дп", PM: "пп"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "дп", PM: "пп"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "дп", PM: "пп"},
			},
		},
	}
}
