package ne

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "y MMMM d, EEEE", Long: "y MMMM d", Medium: "y MMM d", Short: "y-MM-dd"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "जनवरी", Feb: "फेब्रुअरी", Mar: "मार्च", Apr: "अप्रिल", May: "मे", Jun: "जुन", Jul: "जुलाई", Aug: "अगस्ट", Sep: "सेप्टेम्बर", Oct: "अक्टोबर", Nov: "नोभेम्बर", Dec: "डिसेम्बर"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "१", Feb: "२", Mar: "३", Apr: "४", May: "५", Jun: "६", Jul: "७", Aug: "८", Sep: "९", Oct: "१०", Nov: "११", Dec: "१२"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "जनवरी", Feb: "फेब्रुअरी", Mar: "मार्च", Apr: "अप्रिल", May: "मे", Jun: "जुन", Jul: "जुलाई", Aug: "अगस्ट", Sep: "सेप्टेम्बर", Oct: "अक्टोबर", Nov: "नोभेम्बर", Dec: "डिसेम्बर"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "आइत", Mon: "सोम", Tue: "मङ्गल", Wed: "बुध", Thu: "बिही", Fri: "शुक्र", Sat: "शनि"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "आ", Mon: "सो", Tue: "म", Wed: "बु", Thu: "बि", Fri: "शु", Sat: "श"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "आइत", Mon: "सोम", Tue: "मङ्गल", Wed: "बुध", Thu: "बिही", Fri: "शुक्र", Sat: "शनि"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "आइतबार", Mon: "सोमबार", Tue: "मङ्गलबार", Wed: "बुधबार", Thu: "बिहिबार", Fri: "शुक्रबार", Sat: "शनिबार"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "पूर्वाह्न", PM: "अपराह्न"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "पूर्वाह्न", PM: "अपराह्न"},
			},
		},
	}
}
