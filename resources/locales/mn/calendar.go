package mn

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, y 'оны' MM 'сарын' d", Long: "y 'оны' MM 'сарын' d", Medium: "y MMM d", Short: "y-MM-dd"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1}, {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "1-р сар", Feb: "2-р сар", Mar: "3-р сар", Apr: "4-р сар", May: "5-р сар", Jun: "6-р сар", Jul: "7-р сар", Aug: "8-р сар", Sep: "9-р сар", Oct: "10-р сар", Nov: "11-р сар", Dec: "12-р сар"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Нэгдүгээр сар", Feb: "Хоёрдугаар сар", Mar: "Гуравдугаар сар", Apr: "Дөрөвдүгээр сар", May: "Тавдугаар сар", Jun: "Зургадугаар сар", Jul: "Долдугаар сар", Aug: "Наймдугаар сар", Sep: "Есдүгээр сар", Oct: "Аравдугаар сар", Nov: "Арван нэгдүгээр сар", Dec: "Арван хоёрдугаар сар"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Ня", Mon: "Да", Tue: "Мя", Wed: "Лх", Thu: "Пү", Fri: "Ба", Sat: "Бя"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "1", Mon: "2", Tue: "3", Wed: "4", Thu: "5", Fri: "6", Sat: "7"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "Ня", Mon: "Да", Tue: "Мя", Wed: "Лх", Thu: "Пү", Fri: "Ба", Sat: "Бя"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "ням", Mon: "даваа", Tue: "мягмар", Wed: "лхагва", Thu: "пүрэв", Fri: "баасан", Sat: "бямба"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "үө", PM: "үх"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "ҮӨ", PM: "ҮХ"},
			},
		},
	}
}
