package cgg

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "KBZ", Feb: "KBR", Mar: "KST", Apr: "KKN", May: "KTN", Jun: "KMK", Jul: "KMS", Aug: "KMN", Sep: "KMW", Oct: "KKM", Nov: "KNK", Dec: "KNB"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Okwokubanza", Feb: "Okwakabiri", Mar: "Okwakashatu", Apr: "Okwakana", May: "Okwakataana", Jun: "Okwamukaaga", Jul: "Okwamushanju", Aug: "Okwamunaana", Sep: "Okwamwenda", Oct: "Okwaikumi", Nov: "Okwaikumi na kumwe", Dec: "Okwaikumi na ibiri"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "SAN", Mon: "ORK", Tue: "OKB", Wed: "OKS", Thu: "OKN", Fri: "OKT", Sat: "OMK"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "S", Mon: "K", Tue: "R", Wed: "S", Thu: "N", Fri: "T", Sat: "M"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Sande", Mon: "Orwokubanza", Tue: "Orwakabiri", Wed: "Orwakashatu", Thu: "Orwakana", Fri: "Orwakataano", Sat: "Orwamukaaga"},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
