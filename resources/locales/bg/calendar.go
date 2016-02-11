package bg

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y 'г'.", Long: "d MMMM y 'г'.", Medium: "d.MM.y 'г'.", Short: "d.MM.yy 'г'."},
			Time:     ut.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "яну", Feb: "фев", Mar: "март", Apr: "апр", May: "май", Jun: "юни", Jul: "юли", Aug: "авг", Sep: "сеп", Oct: "окт", Nov: "ное", Dec: "дек"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "я", Feb: "ф", Mar: "м", Apr: "а", May: "м", Jun: "ю", Jul: "ю", Aug: "а", Sep: "с", Oct: "о", Nov: "н", Dec: "д"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "януари", Feb: "февруари", Mar: "март", Apr: "април", May: "май", Jun: "юни", Jul: "юли", Aug: "август", Sep: "септември", Oct: "октомври", Nov: "ноември", Dec: "декември"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "нд", Mon: "пн", Tue: "вт", Wed: "ср", Thu: "чт", Fri: "пт", Sat: "сб"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "н", Mon: "п", Tue: "в", Wed: "с", Thu: "ч", Fri: "п", Sat: "с"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "нд", Mon: "пн", Tue: "вт", Wed: "ср", Thu: "чт", Fri: "пт", Sat: "сб"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "неделя", Mon: "понеделник", Tue: "вторник", Wed: "сряда", Thu: "четвъртък", Fri: "петък", Sat: "събота"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "пр.об.", PM: "сл.об."},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "пр.об.", PM: "сл.об."},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "пр.об.", PM: "сл.об."},
			},
		},
	}
}
