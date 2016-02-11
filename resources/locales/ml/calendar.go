package ml

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "y, MMMM d, EEEE", Long: "y, MMMM d", Medium: "y, MMM d", Short: "d/M/yy"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "ജനു", Feb: "ഫെബ്രു", Mar: "മാർ", Apr: "ഏപ്രി", May: "മേയ്", Jun: "ജൂൺ", Jul: "ജൂലൈ", Aug: "ഓഗ", Sep: "സെപ്റ്റം", Oct: "ഒക്ടോ", Nov: "നവം", Dec: "ഡിസം"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "ജ", Feb: "ഫെ", Mar: "മാ", Apr: "ഏ", May: "മെ", Jun: "ജൂ", Jul: "ജൂ", Aug: "ഓ", Sep: "സെ", Oct: "ഒ", Nov: "ന", Dec: "ഡി"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "ജനുവരി", Feb: "ഫെബ്രുവരി", Mar: "മാർച്ച്", Apr: "ഏപ്രിൽ", May: "മേയ്", Jun: "ജൂൺ", Jul: "ജൂലൈ", Aug: "ഓഗസ്റ്റ്", Sep: "സെപ്റ്റംബർ", Oct: "ഒക്\u200cടോബർ", Nov: "നവംബർ", Dec: "ഡിസംബർ"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "ഞായർ", Mon: "തിങ്കൾ", Tue: "ചൊവ്വ", Wed: "ബുധൻ", Thu: "വ്യാഴം", Fri: "വെള്ളി", Sat: "ശനി"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "ഞാ", Mon: "തി", Tue: "ചൊ", Wed: "ബു", Thu: "വ്യാ", Fri: "വെ", Sat: "ശ"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "ഞാ", Mon: "തി", Tue: "ചൊ", Wed: "ബു", Thu: "വ്യാ", Fri: "വെ", Sat: "ശ"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "ഞായറാഴ്\u200cച", Mon: "തിങ്കളാഴ്\u200cച", Tue: "ചൊവ്വാഴ്\u200cച", Wed: "ബുധനാഴ്\u200cച", Thu: "വ്യാഴാഴ്\u200cച", Fri: "വെള്ളിയാഴ്\u200cച", Sat: "ശനിയാഴ്\u200cച"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			},
		},
	}
}
