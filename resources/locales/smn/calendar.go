package smn

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "uđđâivemáánu", Feb: "kuovâmáánu", Mar: "njuhčâmáánu", Apr: "cuáŋuimáánu", May: "vyesimáánu", Jun: "kesimáánu", Jul: "syeinimáánu", Aug: "porgemáánu", Sep: "čohčâmáánu", Oct: "roovvâdmáánu", Nov: "skammâmáánu", Dec: "juovlâmáánu"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "pa", Mon: "vu", Tue: "ma", Wed: "ko", Thu: "tu", Fri: "vá", Sat: "lá"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "P", Mon: "V", Tue: "M", Wed: "K", Thu: "T", Fri: "V", Sat: "L"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "pasepeivi", Mon: "vuossargâ", Tue: "majebargâ", Wed: "koskokko", Thu: "tuorâstâh", Fri: "vástuppeivi", Sat: "lávurdâh"},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
