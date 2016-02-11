package vi

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, 'ngày' dd MMMM 'năm' y", Long: "'Ngày' dd 'tháng' MM 'năm' y", Medium: "d MMM, y", Short: "dd/MM/y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{0} {1}", Long: "{0} {1}", Medium: "{0}, {1}", Short: "{0}, {1}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Thg 1", Feb: "Thg 2", Mar: "Thg 3", Apr: "Thg 4", May: "Thg 5", Jun: "Thg 6", Jul: "Thg 7", Aug: "Thg 8", Sep: "Thg 9", Oct: "Thg 10", Nov: "Thg 11", Dec: "Thg 12"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Tháng 1", Feb: "Tháng 2", Mar: "Tháng 3", Apr: "Tháng 4", May: "Tháng 5", Jun: "Tháng 6", Jul: "Tháng 7", Aug: "Tháng 8", Sep: "Tháng 9", Oct: "Tháng 10", Nov: "Tháng 11", Dec: "Tháng 12"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "CN", Mon: "Th 2", Tue: "Th 3", Wed: "Th 4", Thu: "Th 5", Fri: "Th 6", Sat: "Th 7"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "CN", Mon: "T2", Tue: "T3", Wed: "T4", Thu: "T5", Fri: "T6", Sat: "T7"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "CN", Mon: "T2", Tue: "T3", Wed: "T4", Thu: "T5", Fri: "T6", Sat: "T7"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Chủ Nhật", Mon: "Thứ Hai", Tue: "Thứ Ba", Wed: "Thứ Tư", Thu: "Thứ Năm", Fri: "Thứ Sáu", Sat: "Thứ Bảy"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "SA", PM: "CH"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "SA", PM: "CH"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "SA", PM: "CH"},
			},
		},
	}
}
