package yo

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Ṣẹ́rẹ́", Feb: "Èrèlè", Mar: "Ẹrẹ̀nà", Apr: "Ìgbé", May: "Ẹ̀bibi", Jun: "Òkúdu", Jul: "Agẹmọ", Aug: "Ògún", Sep: "Owewe", Oct: "Ọ̀wàrà", Nov: "Bélú", Dec: "Ọ̀pẹ̀"},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Oṣù Ṣẹ́rẹ́", Feb: "Oṣù Èrèlè", Mar: "Oṣù Ẹrẹ̀nà", Apr: "Oṣù Ìgbé", May: "Oṣù Ẹ̀bibi", Jun: "Oṣù Òkúdu", Jul: "Oṣù Agẹmọ", Aug: "Oṣù Ògún", Sep: "Oṣù Owewe", Oct: "Oṣù Ọ̀wàrà", Nov: "Oṣù Bélú", Dec: "Oṣù Ọ̀pẹ̀"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Àìkú", Mon: "Ajé", Tue: "Ìsẹ́gun", Wed: "Ọjọ́rú", Thu: "Ọjọ́bọ", Fri: "Ẹtì", Sat: "Àbámẹ́ta"},
				Narrow:      ut.CalendarDayFormatNameValue{},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Ọjọ́ Àìkú", Mon: "Ọjọ́ Ajé", Tue: "Ọjọ́ Ìsẹ́gun", Wed: "Ọjọ́rú", Thu: "Ọjọ́bọ", Fri: "Ọjọ́ Ẹtì", Sat: "Ọjọ́ Àbámẹ́ta"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "Àárọ̀", PM: "Ọ̀sán"},
			},
		},
	}
}
