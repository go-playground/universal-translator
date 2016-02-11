package cu

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM 'л'. y.", Long: "y MMMM d", Medium: "y MMM d", Short: "y.MM.dd"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "і҆аⷩ҇", Feb: "феⷡ҇", Mar: "маⷬ҇", Apr: "а҆пⷬ҇", May: "маꙵ", Jun: "і҆ꙋⷩ҇", Jul: "і҆ꙋⷧ҇", Aug: "а҆́ѵⷢ҇", Sep: "сеⷫ҇", Oct: "ѻ҆кⷮ", Nov: "ноеⷨ", Dec: "деⷦ҇"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "І҆", Feb: "Ф", Mar: "М", Apr: "А҆", May: "М", Jun: "І҆", Jul: "І҆", Aug: "А҆", Sep: "С", Oct: "Ѻ҆", Nov: "Н", Dec: "Д"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "і҆аннꙋа́рїй", Feb: "феврꙋа́рїй", Mar: "ма́ртъ", Apr: "а҆прі́ллїй", May: "ма́їй", Jun: "і҆ꙋ́нїй", Jul: "і҆ꙋ́лїй", Aug: "а҆́ѵгꙋстъ", Sep: "септе́мврїй", Oct: "ѻ҆ктѡ́врїй", Nov: "ное́мврїй", Dec: "деке́мврїй"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "ндⷧ҇ѧ", Mon: "пнⷣе", Tue: "втоⷬ҇", Wed: "срⷣе", Thu: "чеⷦ҇", Fri: "пѧⷦ҇", Sat: "сꙋⷠ҇"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "Н", Mon: "П", Tue: "В", Wed: "С", Thu: "Ч", Fri: "П", Sat: "С"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "ндⷧ҇ѧ", Mon: "пнⷣе", Tue: "втоⷬ҇", Wed: "срⷣе", Thu: "чеⷦ҇", Fri: "пѧⷦ҇", Sat: "сꙋⷠ҇"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "недѣ́лѧ", Mon: "понедѣ́льникъ", Tue: "вто́рникъ", Wed: "среда̀", Thu: "четверто́къ", Fri: "пѧто́къ", Sat: "сꙋббѡ́та"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "ДП", PM: "ПП"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "ДП", PM: "ПП"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "ДП", PM: "ПП"},
			},
		},
	}
}
