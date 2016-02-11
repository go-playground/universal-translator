package as

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM, y", Medium: "dd-MM-y", Short: "d-M-y"},
			Time:     ut.CalendarDateFormat{Full: "h.mm.ss a zzzz", Long: "h.mm.ss a z", Medium: "h.mm.ss a", Short: "h.mm. a"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "জানু", Feb: "ফেব্ৰু", Mar: "মাৰ্চ", Apr: "এপ্ৰিল", May: "মে", Jun: "জুন", Jul: "জুলাই", Aug: "আগ", Sep: "সেপ্ট", Oct: "অক্টো", Nov: "নভে", Dec: "ডিসে"},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "জানুৱাৰী", Feb: "ফেব্ৰুৱাৰী", Mar: "মাৰ্চ", Apr: "এপ্ৰিল", May: "মে", Jun: "জুন", Jul: "জুলাই", Aug: "আগষ্ট", Sep: "ছেপ্তেম্বৰ", Oct: "অক্টোবৰ", Nov: "নৱেম্বৰ", Dec: "ডিচেম্বৰ"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "ৰবি", Mon: "সোম", Tue: "মঙ্গল", Wed: "বুধ", Thu: "বৃহষ্পতি", Fri: "শুক্ৰ", Sat: "শনি"},
				Narrow:      ut.CalendarDayFormatNameValue{},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "দেওবাৰ", Mon: "সোমবাৰ", Tue: "মঙ্গলবাৰ", Wed: "বুধবাৰ", Thu: "বৃহষ্পতিবাৰ", Fri: "শুক্ৰবাৰ", Sat: "শনিবাৰ"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "পূৰ্বাহ্ণ", PM: "অপৰাহ্ণ"},
			},
		},
	}
}
