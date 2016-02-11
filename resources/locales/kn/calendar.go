package kn

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "d/M/yy"},
			Time:     ut.CalendarDateFormat{Full: "hh:mm:ss a zzzz", Long: "hh:mm:ss a z", Medium: "hh:mm:ss a", Short: "hh:mm a"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "ಜನ", Feb: "ಫೆಬ್ರ", Mar: "ಮಾರ್ಚ್", Apr: "ಏಪ್ರಿ", May: "ಮೇ", Jun: "ಜೂನ್", Jul: "ಜುಲೈ", Aug: "ಆಗ", Sep: "ಸೆಪ್ಟೆಂ", Oct: "ಅಕ್ಟೋ", Nov: "ನವೆಂ", Dec: "ಡಿಸೆಂ"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "ಜ", Feb: "ಫೆ", Mar: "ಮಾ", Apr: "ಏ", May: "ಮೇ", Jun: "ಜೂ", Jul: "ಜು", Aug: "ಆ", Sep: "ಸೆ", Oct: "ಅ", Nov: "ನ", Dec: "ಡಿ"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "ಜನವರಿ", Feb: "ಫೆಬ್ರವರಿ", Mar: "ಮಾರ್ಚ್", Apr: "ಏಪ್ರಿಲ್", May: "ಮೇ", Jun: "ಜೂನ್", Jul: "ಜುಲೈ", Aug: "ಆಗಸ್ಟ್", Sep: "ಸೆಪ್ಟೆಂಬರ್", Oct: "ಅಕ್ಟೋಬರ್", Nov: "ನವೆಂಬರ್", Dec: "ಡಿಸೆಂಬರ್"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "ಭಾನು", Mon: "ಸೋಮ", Tue: "ಮಂಗಳ", Wed: "ಬುಧ", Thu: "ಗುರು", Fri: "ಶುಕ್ರ", Sat: "ಶನಿ"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "ಭಾ", Mon: "ಸೋ", Tue: "ಮಂ", Wed: "ಬು", Thu: "ಗು", Fri: "ಶು", Sat: "ಶ"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "ಭಾನು", Mon: "ಸೋಮ", Tue: "ಮಂಗಳ", Wed: "ಬುಧ", Thu: "ಗುರು", Fri: "ಶುಕ್ರ", Sat: "ಶನಿ"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "ಭಾನುವಾರ", Mon: "ಸೋಮವಾರ", Tue: "ಮಂಗಳವಾರ", Wed: "ಬುಧವಾರ", Thu: "ಗುರುವಾರ", Fri: "ಶುಕ್ರವಾರ", Sat: "ಶನಿವಾರ"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "ಪೂರ್ವಾಹ್ನ", PM: "ಅಪರಾಹ್ನ"},
			},
		},
	}
}
