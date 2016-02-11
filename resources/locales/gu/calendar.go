package gu

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "d/M/yy"},
			Time:     ut.CalendarDateFormat{Full: "hh:mm:ss a zzzz", Long: "hh:mm:ss a z", Medium: "hh:mm:ss a", Short: "hh:mm a"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "જાન્યુ", Feb: "ફેબ્રુ", Mar: "માર્ચ", Apr: "એપ્રિલ", May: "મે", Jun: "જૂન", Jul: "જુલાઈ", Aug: "ઑગસ્ટ", Sep: "સપ્ટે", Oct: "ઑક્ટો", Nov: "નવે", Dec: "ડિસે"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "જા", Feb: "ફે", Mar: "મા", Apr: "એ", May: "મે", Jun: "જૂ", Jul: "જુ", Aug: "ઑ", Sep: "સ", Oct: "ઑ", Nov: "ન", Dec: "ડિ"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "જાન્યુઆરી", Feb: "ફેબ્રુઆરી", Mar: "માર્ચ", Apr: "એપ્રિલ", May: "મે", Jun: "જૂન", Jul: "જુલાઈ", Aug: "ઑગસ્ટ", Sep: "સપ્ટેમ્બર", Oct: "ઑક્ટોબર", Nov: "નવેમ્બર", Dec: "ડિસેમ્બર"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "રવિ", Mon: "સોમ", Tue: "મંગળ", Wed: "બુધ", Thu: "ગુરુ", Fri: "શુક્ર", Sat: "શનિ"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "ર", Mon: "સો", Tue: "મં", Wed: "બુ", Thu: "ગુ", Fri: "શુ", Sat: "શ"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "ર", Mon: "સો", Tue: "મં", Wed: "બુ", Thu: "ગુ", Fri: "શુ", Sat: "શ"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "રવિવાર", Mon: "સોમવાર", Tue: "મંગળવાર", Wed: "બુધવાર", Thu: "ગુરુવાર", Fri: "શુક્રવાર", Sat: "શનિવાર"},
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
