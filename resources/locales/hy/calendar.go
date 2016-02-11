package hy

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "yթ. MMMM d, EEEE", Long: "dd MMMM, yթ.", Medium: "dd MMM, yթ.", Short: "dd.MM.yy"},
			Time:     ut.CalendarDateFormat{Full: "H:mm:ss, zzzz", Long: "H:mm:ss, z", Medium: "H:mm:ss", Short: "H:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "հնվ", Feb: "փտվ", Mar: "մրտ", Apr: "ապր", May: "մյս", Jun: "հնս", Jul: "հլս", Aug: "օգս", Sep: "սեպ", Oct: "հոկ", Nov: "նոյ", Dec: "դեկ"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "Հ", Feb: "Փ", Mar: "Մ", Apr: "Ա", May: "Մ", Jun: "Հ", Jul: "Հ", Aug: "Օ", Sep: "Ս", Oct: "Հ", Nov: "Ն", Dec: "Դ"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "հունվար", Feb: "փետրվար", Mar: "մարտ", Apr: "ապրիլ", May: "մայիս", Jun: "հունիս", Jul: "հուլիս", Aug: "օգոստոս", Sep: "սեպտեմբեր", Oct: "հոկտեմբեր", Nov: "նոյեմբեր", Dec: "դեկտեմբեր"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "կիր", Mon: "երկ", Tue: "երք", Wed: "չրք", Thu: "հնգ", Fri: "ուր", Sat: "շբթ"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "Կր", Mon: "Եկ", Tue: "Եր", Wed: "Չր", Thu: "Հգ", Fri: "Ու", Sat: "Շբ"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "կիր", Mon: "երկ", Tue: "երք", Wed: "չրք", Thu: "հնգ", Fri: "ուր", Sat: "շբթ"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "կիրակի", Mon: "երկուշաբթի", Tue: "երեքշաբթի", Wed: "չորեքշաբթի", Thu: "հինգշաբթի", Fri: "ուրբաթ", Sat: "շաբաթ"},
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
