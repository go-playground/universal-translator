package kk

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "y 'ж'. d MMMM, EEEE", Long: "y 'ж'. d MMMM", Medium: "y 'ж'. dd MMM", Short: "dd.MM.yy"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Қаң.", Feb: "Ақп.", Mar: "Нау.", Apr: "Сәу.", May: "Мам.", Jun: "Мау.", Jul: "Шіл.", Aug: "Там.", Sep: "Қыр.", Oct: "Қаз.", Nov: "Қар.", Dec: "Жел."},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "Қ", Feb: "А", Mar: "Н", Apr: "С", May: "М", Jun: "М", Jul: "Ш", Aug: "Т", Sep: "Қ", Oct: "Қ", Nov: "Қ", Dec: "Ж"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Қаңтар", Feb: "Ақпан", Mar: "Наурыз", Apr: "Сәуір", May: "Мамыр", Jun: "Маусым", Jul: "Шілде", Aug: "Тамыз", Sep: "Қыркүйек", Oct: "Қазан", Nov: "Қараша", Dec: "Желтоқсан"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Жс", Mon: "Дс", Tue: "Сс", Wed: "Ср", Thu: "Бс", Fri: "Жм", Sat: "Сб"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "Ж", Mon: "Д", Tue: "С", Wed: "С", Thu: "Б", Fri: "Ж", Sat: "С"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "Жс", Mon: "Дс", Tue: "Сс", Wed: "Ср", Thu: "Бс", Fri: "Жм", Sat: "Сб"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Жексенбі", Mon: "Дүйсенбі", Tue: "Сейсенбі", Wed: "Сәрсенбі", Thu: "Бейсенбі", Fri: "Жұма", Sat: "Сенбі"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "таң", PM: "түс/кеш"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "таң", PM: "түс/кеш"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "таң", PM: "түс/кеш"},
			},
		},
	}
}
