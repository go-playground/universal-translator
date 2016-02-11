package os

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM, y 'аз'", Long: "d MMMM, y 'аз'", Medium: "dd MMM y 'аз'", Short: "dd.MM.yy"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Янв.", Feb: "Февр.", Mar: "Март.", Apr: "Апр.", May: "Май", Jun: "Июнь", Jul: "Июль", Aug: "Авг.", Sep: "Сент.", Oct: "Окт.", Nov: "Нояб.", Dec: "Дек."},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "Я", Feb: "Ф", Mar: "М", Apr: "А", May: "М", Jun: "И", Jul: "И", Aug: "А", Sep: "С", Oct: "О", Nov: "Н", Dec: "Д"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Январь", Feb: "Февраль", Mar: "Мартъи", Apr: "Апрель", May: "Май", Jun: "Июнь", Jul: "Июль", Aug: "Август", Sep: "Сентябрь", Oct: "Октябрь", Nov: "Ноябрь", Dec: "Декабрь"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Хцб", Mon: "Крс", Tue: "Дцг", Wed: "Ӕрт", Thu: "Цпр", Fri: "Мрб", Sat: "Сбт"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "Х", Mon: "К", Tue: "Д", Wed: "Ӕ", Thu: "Ц", Fri: "М", Sat: "С"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Хуыцаубон", Mon: "Къуырисӕр", Tue: "Дыццӕг", Wed: "Ӕртыццӕг", Thu: "Цыппӕрӕм", Fri: "Майрӕмбон", Sat: "Сабат"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "ӕмбисбоны размӕ", PM: "ӕмбисбоны фӕстӕ"},
			},
		},
	}
}
