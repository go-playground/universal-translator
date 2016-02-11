package hi

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "dd/MM/y", Short: "d/M/yy"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{Full: "{1} को {0}", Long: "{1} को {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "जन॰", Feb: "फ़र॰", Mar: "मार्च", Apr: "अप्रैल", May: "मई", Jun: "जून", Jul: "जुल॰", Aug: "अग॰", Sep: "सित॰", Oct: "अक्तू॰", Nov: "नव॰", Dec: "दिस॰"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "ज", Feb: "फ़", Mar: "मा", Apr: "अ", May: "म", Jun: "जू", Jul: "जु", Aug: "अ", Sep: "सि", Oct: "अ", Nov: "न", Dec: "दि"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "जनवरी", Feb: "फ़रवरी", Mar: "मार्च", Apr: "अप्रैल", May: "मई", Jun: "जून", Jul: "जुलाई", Aug: "अगस्त", Sep: "सितंबर", Oct: "अक्तूबर", Nov: "नवंबर", Dec: "दिसंबर"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "रवि", Mon: "सोम", Tue: "मंगल", Wed: "बुध", Thu: "गुरु", Fri: "शुक्र", Sat: "शनि"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "र", Mon: "सो", Tue: "मं", Wed: "बु", Thu: "गु", Fri: "शु", Sat: "श"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "र", Mon: "सो", Tue: "मं", Wed: "बु", Thu: "गु", Fri: "शु", Sat: "श"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "रविवार", Mon: "सोमवार", Tue: "मंगलवार", Wed: "बुधवार", Thu: "गुरुवार", Fri: "शुक्रवार", Sat: "शनिवार"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "पूर्व", PM: "अपर"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "पू", PM: "अ"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "पूर्वाह्न", PM: "अपराह्न"},
			},
		},
	}
}
