package bas

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "kɔn", Feb: "mac", Mar: "mat", Apr: "mto", May: "mpu", Jun: "hil", Jul: "nje", Aug: "hik", Sep: "dip", Oct: "bio", Nov: "may", Dec: "liɓ"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "k", Feb: "m", Mar: "m", Apr: "m", May: "m", Jun: "h", Jul: "n", Aug: "h", Sep: "d", Oct: "b", Nov: "m", Dec: "l"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Kɔndɔŋ", Feb: "Màcɛ̂l", Mar: "Màtùmb", Apr: "Màtop", May: "M̀puyɛ", Jun: "Hìlòndɛ̀", Jul: "Njèbà", Aug: "Hìkaŋ", Sep: "Dìpɔ̀s", Oct: "Bìòôm", Nov: "Màyɛsèp", Dec: "Lìbuy li ńyèe"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "nɔy", Mon: "nja", Tue: "uum", Wed: "ŋge", Thu: "mbɔ", Fri: "kɔɔ", Sat: "jon"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "n", Mon: "n", Tue: "u", Wed: "ŋ", Thu: "m", Fri: "k", Sat: "j"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "ŋgwà nɔ̂y", Mon: "ŋgwà njaŋgumba", Tue: "ŋgwà ûm", Wed: "ŋgwà ŋgê", Thu: "ŋgwà mbɔk", Fri: "ŋgwà kɔɔ", Sat: "ŋgwà jôn"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "I bikɛ̂glà", PM: "I ɓugajɔp"},
			},
		},
	}
}
