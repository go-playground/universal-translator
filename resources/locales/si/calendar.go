package si

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "y MMMM d, EEEE", Long: "y MMMM d", Medium: "y MMM d", Short: "y-MM-dd"},
			Time:     ut.CalendarDateFormat{Full: "HH.mm.ss zzzz", Long: "HH.mm.ss z", Medium: "HH.mm.ss", Short: "HH.mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "ජන", Feb: "පෙබ", Mar: "මාර්", Apr: "අප්\u200dරේල්", May: "මැයි", Jun: "ජූනි", Jul: "ජූලි", Aug: "අගෝ", Sep: "සැප්", Oct: "ඔක්", Nov: "නොවැ", Dec: "දෙසැ"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "ජ", Feb: "පෙ", Mar: "මා", Apr: "අ", May: "මැ", Jun: "ජූ", Jul: "ජූ", Aug: "අ", Sep: "සැ", Oct: "ඔ", Nov: "නෙ", Dec: "දෙ"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "ජනවාරි", Feb: "පෙබරවාරි", Mar: "මාර්තු", Apr: "අප්\u200dරේල්", May: "මැයි", Jun: "ජූනි", Jul: "ජූලි", Aug: "අගෝස්තු", Sep: "සැප්තැම්බර්", Oct: "ඔක්තෝබර්", Nov: "නොවැම්බර්", Dec: "දෙසැම්බර්"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "ඉරිදා", Mon: "සඳුදා", Tue: "අඟහ", Wed: "බදාදා", Thu: "බ්\u200dරහස්", Fri: "සිකු", Sat: "සෙන"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "ඉ", Mon: "ස", Tue: "අ", Wed: "බ", Thu: "බ්\u200dර", Fri: "සි", Sat: "සෙ"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "ඉරි", Mon: "සඳු", Tue: "අඟ", Wed: "බදා", Thu: "බ්\u200dරහ", Fri: "සිකු", Sat: "සෙන"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "ඉරිදා", Mon: "සඳුදා", Tue: "අඟහරුවාදා", Wed: "බදාදා", Thu: "බ්\u200dරහස්පතින්දා", Fri: "සිකුරාදා", Sat: "සෙනසුරාදා"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "පෙ.ව.", PM: "ප.ව."},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "පෙ", PM: "ප"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "පෙ.ව.", PM: "ප.ව."},
			},
		},
	}
}
