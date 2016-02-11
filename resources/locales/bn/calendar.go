package bn

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "d/M/yy"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "জানুয়ারী", Feb: "ফেব্রুয়ারী", Mar: "মার্চ", Apr: "এপ্রিল", May: "মে", Jun: "জুন", Jul: "জুলাই", Aug: "আগস্ট", Sep: "সেপ্টেম্বর", Oct: "অক্টোবর", Nov: "নভেম্বর", Dec: "ডিসেম্বর"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "জা", Feb: "ফে", Mar: "মা", Apr: "এ", May: "মে", Jun: "জুন", Jul: "জু", Aug: "আ", Sep: "সে", Oct: "অ", Nov: "ন", Dec: "ডি"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "জানুয়ারী", Feb: "ফেব্রুয়ারী", Mar: "মার্চ", Apr: "এপ্রিল", May: "মে", Jun: "জুন", Jul: "জুলাই", Aug: "আগস্ট", Sep: "সেপ্টেম্বর", Oct: "অক্টোবর", Nov: "নভেম্বর", Dec: "ডিসেম্বর"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "রবি", Mon: "সোম", Tue: "মঙ্গল", Wed: "বুধ", Thu: "বৃহস্পতি", Fri: "শুক্র", Sat: "শনি"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "র", Mon: "সো", Tue: "ম", Wed: "বু", Thu: "বৃ", Fri: "শু", Sat: "শ"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "রঃ", Mon: "সোঃ", Tue: "মঃ", Wed: "বুঃ", Thu: "বৃঃ", Fri: "শুঃ", Sat: "শনি"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "রবিবার", Mon: "সোমবার", Tue: "মঙ্গলবার", Wed: "বুধবার", Thu: "বৃহষ্পতিবার", Fri: "শুক্রবার", Sat: "শনিবার"},
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
