package sr

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, dd. MMMM y.", Long: "dd. MMMM y.", Medium: "dd.MM.y.", Short: "d.M.yy."},
			Time:     ut.CalendarDateFormat{Full: "HH.mm.ss zzzz", Long: "HH.mm.ss z", Medium: "HH.mm.ss", Short: "HH.mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "јан", Feb: "феб", Mar: "мар", Apr: "апр", May: "мај", Jun: "јун", Jul: "јул", Aug: "авг", Sep: "сеп", Oct: "окт", Nov: "нов", Dec: "дец"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "ј", Feb: "ф", Mar: "м", Apr: "а", May: "м", Jun: "ј", Jul: "ј", Aug: "а", Sep: "с", Oct: "о", Nov: "н", Dec: "д"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "јануар", Feb: "фебруар", Mar: "март", Apr: "април", May: "мај", Jun: "јун", Jul: "јул", Aug: "август", Sep: "септембар", Oct: "октобар", Nov: "новембар", Dec: "децембар"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "нед", Mon: "пон", Tue: "уто", Wed: "сре", Thu: "чет", Fri: "пет", Sat: "суб"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "н", Mon: "п", Tue: "у", Wed: "с", Thu: "ч", Fri: "п", Sat: "с"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "не", Mon: "по", Tue: "ут", Wed: "ср", Thu: "че", Fri: "пе", Sat: "су"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "недеља", Mon: "понедељак", Tue: "уторак", Wed: "среда", Thu: "четвртак", Fri: "петак", Sat: "субота"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "пре подне", PM: "по подне"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "пре подне", PM: "по подне"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "пре подне", PM: "по подне"},
			},
		},
	}
}
