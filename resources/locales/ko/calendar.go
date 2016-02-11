package ko

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "y년 M월 d일 EEEE", Long: "y년 M월 d일", Medium: "y. M. d.", Short: "yy. M. d."},
			Time:     ut.CalendarDateFormat{Full: "a h시 m분 s초 zzzz", Long: "a h시 m분 s초 z", Medium: "a h:mm:ss", Short: "a h:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "1월", Feb: "2월", Mar: "3월", Apr: "4월", May: "5월", Jun: "6월", Jul: "7월", Aug: "8월", Sep: "9월", Oct: "10월", Nov: "11월", Dec: "12월"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "1월", Feb: "2월", Mar: "3월", Apr: "4월", May: "5월", Jun: "6월", Jul: "7월", Aug: "8월", Sep: "9월", Oct: "10월", Nov: "11월", Dec: "12월"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "1월", Feb: "2월", Mar: "3월", Apr: "4월", May: "5월", Jun: "6월", Jul: "7월", Aug: "8월", Sep: "9월", Oct: "10월", Nov: "11월", Dec: "12월"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "일", Mon: "월", Tue: "화", Wed: "수", Thu: "목", Fri: "금", Sat: "토"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "일", Mon: "월", Tue: "화", Wed: "수", Thu: "목", Fri: "금", Sat: "토"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "일", Mon: "월", Tue: "화", Wed: "수", Thu: "목", Fri: "금", Sat: "토"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "일요일", Mon: "월요일", Tue: "화요일", Wed: "수요일", Thu: "목요일", Fri: "금요일", Sat: "토요일"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "오전", PM: "오후"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "오전", PM: "오후"},
			},
		},
	}
}
