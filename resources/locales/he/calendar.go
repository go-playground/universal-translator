package he

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d בMMMM y", Long: "d בMMMM y", Medium: "d בMMM y", Short: "d.M.y"},
			Time:     ut.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} בשעה {0}", Long: "{1} בשעה {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "ינו׳", Feb: "פבר׳", Mar: "מרץ", Apr: "אפר׳", May: "מאי", Jun: "יוני", Jul: "יולי", Aug: "אוג׳", Sep: "ספט׳", Oct: "אוק׳", Nov: "נוב׳", Dec: "דצמ׳"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "ינואר", Feb: "פברואר", Mar: "מרץ", Apr: "אפריל", May: "מאי", Jun: "יוני", Jul: "יולי", Aug: "אוגוסט", Sep: "ספטמבר", Oct: "אוקטובר", Nov: "נובמבר", Dec: "דצמבר"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "יום א׳", Mon: "יום ב׳", Tue: "יום ג׳", Wed: "יום ד׳", Thu: "יום ה׳", Fri: "יום ו׳", Sat: "שבת"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "א׳", Mon: "ב׳", Tue: "ג׳", Wed: "ד׳", Thu: "ה׳", Fri: "ו׳", Sat: "ש׳"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "א׳", Mon: "ב׳", Tue: "ג׳", Wed: "ד׳", Thu: "ה׳", Fri: "ו׳", Sat: "ש׳"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "יום ראשון", Mon: "יום שני", Tue: "יום שלישי", Wed: "יום רביעי", Thu: "יום חמישי", Fri: "יום שישי", Sat: "יום שבת"},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
