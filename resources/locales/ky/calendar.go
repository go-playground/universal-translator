package ky

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d-MMMM, y-'ж'.", Long: "y MMMM d", Medium: "y MMM d", Short: "dd.MM.yy"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Янв", Feb: "Фев", Mar: "Мар", Apr: "Апр", May: "Май", Jun: "Июн", Jul: "Июл", Aug: "Авг", Sep: "Сен", Oct: "Окт", Nov: "Ноя", Dec: "Дек"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "Я", Feb: "Ф", Mar: "М", Apr: "А", May: "М", Jun: "И", Jul: "И", Aug: "А", Sep: "С", Oct: "О", Nov: "Н", Dec: "Д"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Январь", Feb: "Февраль", Mar: "Март", Apr: "Апрель", May: "Май", Jun: "Июнь", Jul: "Июль", Aug: "Август", Sep: "Сентябрь", Oct: "Октябрь", Nov: "Ноябрь", Dec: "Декабрь"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "жек.", Mon: "дүй.", Tue: "шейш.", Wed: "шарш.", Thu: "бейш.", Fri: "жума", Sat: "ишм."},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "Ж", Mon: "Д", Tue: "Ш", Wed: "Ш", Thu: "Б", Fri: "Ж", Sat: "И"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "жк", Mon: "дш.", Tue: "шш.", Wed: "шр.", Thu: "бш.", Fri: "жм.", Sat: "иш."},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "жекшемби", Mon: "дүйшөмбү", Tue: "шейшемби", Wed: "шаршемби", Thu: "бейшемби", Fri: "жума", Sat: "ишемби"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "тң", PM: "тк"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "тң", PM: "тк"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "таңкы", PM: "түштөн кийинки"},
			},
		},
	}
}
