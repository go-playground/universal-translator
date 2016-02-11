package lrc

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "y MMMM d, EEEE", Long: "y MMMM d", Medium: "y MMM d", Short: "y-MM-dd"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "جانڤیە", Feb: "فئڤریە", Mar: "مارس", Apr: "آڤریل", May: "مئی", Jun: "جوٙأن", Jul: "جوٙلا", Aug: "آگوست", Sep: "سئپتامر", Oct: "ئوکتوڤر", Nov: "نوڤامر", Dec: "دئسامر"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "جانڤیە", Feb: "فئڤریە", Mar: "مارس", Apr: "آڤریل", May: "مئی", Jun: "جوٙأن", Jul: "جوٙلا", Aug: "آگوست", Sep: "سئپتامر", Oct: "ئوکتوڤر", Nov: "نوڤامر", Dec: "دئسامر"},
			},
			Days: ut.CalendarDayFormatNames{},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			},
		},
	}
}
