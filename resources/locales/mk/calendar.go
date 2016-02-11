package mk

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, dd MMMM y", Long: "dd MMMM y", Medium: "dd.M.y", Short: "dd.M.yy"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "јан.", Feb: "фев.", Mar: "мар.", Apr: "апр.", May: "мај", Jun: "јун.", Jul: "јул.", Aug: "авг.", Sep: "септ.", Oct: "окт.", Nov: "ноем.", Dec: "дек."},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "ј", Feb: "ф", Mar: "м", Apr: "а", May: "м", Jun: "ј", Jul: "ј", Aug: "а", Sep: "с", Oct: "о", Nov: "н", Dec: "д"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "јануари", Feb: "февруари", Mar: "март", Apr: "април", May: "мај", Jun: "јуни", Jul: "јули", Aug: "август", Sep: "септември", Oct: "октомври", Nov: "ноември", Dec: "декември"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "нед.", Mon: "пон.", Tue: "вт.", Wed: "сре.", Thu: "чет.", Fri: "пет.", Sat: "саб."},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "н", Mon: "п", Tue: "в", Wed: "с", Thu: "ч", Fri: "п", Sat: "с"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "нед.", Mon: "пон.", Tue: "вто.", Wed: "сре.", Thu: "чет.", Fri: "пет.", Sat: "саб."},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "недела", Mon: "понеделник", Tue: "вторник", Wed: "среда", Thu: "четврток", Fri: "петок", Sat: "сабота"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "претпладне", PM: "попладне"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "прет.", PM: "попл."},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "претпладне", PM: "попладне"},
			},
		},
	}
}
