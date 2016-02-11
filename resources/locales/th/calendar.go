package th

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEEที่ d MMMM G y", Long: "d MMMM G y", Medium: "d MMM y", Short: "d/M/yy"},
			Time:     ut.CalendarDateFormat{Full: "H นาฬิกา mm นาที ss วินาที zzzz", Long: "H นาฬิกา mm นาที ss วินาที z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "ม.ค.", Feb: "ก.พ.", Mar: "มี.ค.", Apr: "เม.ย.", May: "พ.ค.", Jun: "มิ.ย.", Jul: "ก.ค.", Aug: "ส.ค.", Sep: "ก.ย.", Oct: "ต.ค.", Nov: "พ.ย.", Dec: "ธ.ค."},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "ม.ค.", Feb: "ก.พ.", Mar: "มี.ค.", Apr: "เม.ย.", May: "พ.ค.", Jun: "มิ.ย.", Jul: "ก.ค.", Aug: "ส.ค.", Sep: "ก.ย.", Oct: "ต.ค.", Nov: "พ.ย.", Dec: "ธ.ค."},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "มกราคม", Feb: "กุมภาพันธ์", Mar: "มีนาคม", Apr: "เมษายน", May: "พฤษภาคม", Jun: "มิถุนายน", Jul: "กรกฎาคม", Aug: "สิงหาคม", Sep: "กันยายน", Oct: "ตุลาคม", Nov: "พฤศจิกายน", Dec: "ธันวาคม"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "อา.", Mon: "จ.", Tue: "อ.", Wed: "พ.", Thu: "พฤ.", Fri: "ศ.", Sat: "ส."},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "อา", Mon: "จ", Tue: "อ", Wed: "พ", Thu: "พฤ", Fri: "ศ", Sat: "ส"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "อา.", Mon: "จ.", Tue: "อ.", Wed: "พ.", Thu: "พฤ.", Fri: "ศ.", Sat: "ส."},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "วันอาทิตย์", Mon: "วันจันทร์", Tue: "วันอังคาร", Wed: "วันพุธ", Thu: "วันพฤหัสบดี", Fri: "วันศุกร์", Sat: "วันเสาร์"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "ก่อนเที่ยง", PM: "หลังเที่ยง"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "ก่อนเที่ยง", PM: "หลังเที่ยง"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "ก่อนเที่ยง", PM: "หลังเที่ยง"},
			},
		},
	}
}
