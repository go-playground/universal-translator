package sah

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "y 'сыл' MMMM d 'күнэ', EEEE", Long: "y, MMMM d", Medium: "y, MMM d", Short: "yy/M/d"},
			Time:     ut.CalendarDateFormat{Full: "", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: ""},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Тохс", Feb: "Олун", Mar: "Клн_ттр", Apr: "Мус_уст", May: "Ыам_йн", Jun: "Бэс_йн", Jul: "От_йн", Aug: "Атрдь_йн", Sep: "Блҕн_йн", Oct: "Алт", Nov: "Сэт", Dec: "Ахс"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "Т", Feb: "О", Mar: "К", Apr: "М", May: "Ы", Jun: "Б", Jul: "О", Aug: "А", Sep: "Б", Oct: "А", Nov: "С", Dec: "А"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Тохсунньу", Feb: "Олунньу", Mar: "Кулун тутар", Apr: "Муус устар", May: "Ыам ыйын", Jun: "Бэс ыйын", Jul: "От ыйын", Aug: "Атырдьых ыйын", Sep: "Балаҕан ыйын", Oct: "Алтынньы", Nov: "Сэтинньи", Dec: "Ахсынньы"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Бс", Mon: "Бн", Tue: "Оп", Wed: "Сэ", Thu: "Чп", Fri: "Бэ", Sat: "Сб"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "Б", Mon: "Б", Tue: "О", Wed: "С", Thu: "Ч", Fri: "Б", Sat: "С"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Баскыһыанньа", Mon: "Бэнидиэлинньик", Tue: "Оптуорунньук", Wed: "Сэрэдэ", Thu: "Чэппиэр", Fri: "Бээтиҥсэ", Sat: "Субуота"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "ЭИ", PM: "ЭК"},
			},
		},
	}
}
