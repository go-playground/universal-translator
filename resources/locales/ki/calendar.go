package ki

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "JEN", Feb: "WKR", Mar: "WGT", Apr: "WKN", May: "WTN", Jun: "WTD", Jul: "WMJ", Aug: "WNN", Sep: "WKD", Oct: "WIK", Nov: "WMW", Dec: "DIT"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "K", Mar: "G", Apr: "K", May: "G", Jun: "G", Jul: "M", Aug: "K", Sep: "K", Oct: "I", Nov: "I", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Njenuarĩ", Feb: "Mwere wa kerĩ", Mar: "Mwere wa gatatũ", Apr: "Mwere wa kana", May: "Mwere wa gatano", Jun: "Mwere wa gatandatũ", Jul: "Mwere wa mũgwanja", Aug: "Mwere wa kanana", Sep: "Mwere wa kenda", Oct: "Mwere wa ikũmi", Nov: "Mwere wa ikũmi na ũmwe", Dec: "Ndithemba"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "KMA", Mon: "NTT", Tue: "NMN", Wed: "NMT", Thu: "ART", Fri: "NMA", Sat: "NMM"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "K", Mon: "N", Tue: "N", Wed: "N", Thu: "A", Fri: "N", Sat: "N"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Kiumia", Mon: "Njumatatũ", Tue: "Njumaine", Wed: "Njumatana", Thu: "Aramithi", Fri: "Njumaa", Sat: "Njumamothi"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "Kiroko", PM: "Hwaĩ-inĩ"},
			},
		},
	}
}
