package brx

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "ज", Feb: "फे", Mar: "मा", Apr: "ए", May: "मे", Jun: "जु", Jul: "जु", Aug: "आ", Sep: "से", Oct: "अ", Nov: "न", Dec: "दि"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "जानुवारी", Feb: "फेब्रुवारी", Mar: "मार्स", Apr: "एफ्रिल", May: "मे", Jun: "जुन", Jul: "जुलाइ", Aug: "आगस्थ", Sep: "सेबथेज्ब़र", Oct: "अखथबर", Nov: "नबेज्ब़र", Dec: "दिसेज्ब़र"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "रबि", Mon: "सम", Tue: "मंगल", Wed: "बुद", Thu: "बिसथि", Fri: "सुखुर", Sat: "सुनि"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "र", Mon: "स", Tue: "मं", Wed: "बु", Thu: "बि", Fri: "सु", Sat: "सु"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "रबिबार", Mon: "समबार", Tue: "मंगलबार", Wed: "बुदबार", Thu: "बिसथिबार", Fri: "सुखुरबार", Sat: "सुनिबार"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "फुं", PM: "बेलासे"},
			},
		},
	}
}
