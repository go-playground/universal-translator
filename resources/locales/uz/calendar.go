package uz

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, y MMMM dd", Long: "d-MMMM, y", Medium: "d-MMM, y", Short: "dd/MM/yy"},
			Time:     ut.CalendarDateFormat{Full: "H:mm:ss (zzzz)", Long: "H:mm:ss (z)", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Yanv", Feb: "Fev", Mar: "Mar", Apr: "Apr", May: "May", Jun: "Iyun", Jul: "Iyul", Aug: "Avg", Sep: "Sen", Oct: "Okt", Nov: "Noya", Dec: "Dek"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "Y", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "I", Jul: "I", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Yanvar", Feb: "Fevral", Mar: "Mart", Apr: "Aprel", May: "May", Jun: "Iyun", Jul: "Iyul", Aug: "Avgust", Sep: "Sentabr", Oct: "Oktabr", Nov: "Noyabr", Dec: "Dekabr"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Ya", Mon: "Du", Tue: "Se", Wed: "Ch", Thu: "Pa", Fri: "Ju", Sat: "Sh"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "Y", Mon: "D", Tue: "S", Wed: "C", Thu: "P", Fri: "J", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "Ya", Mon: "Du", Tue: "Se", Wed: "Ch", Thu: "Pa", Fri: "Ju", Sat: "Sh"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "yakshanba", Mon: "dushanba", Tue: "seshanba", Wed: "chorshanba", Thu: "payshanba", Fri: "juma", Sat: "shanba"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "TO", PM: "TK"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "TO", PM: "TK"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "TO", PM: "TK"},
			},
		},
	}
}
