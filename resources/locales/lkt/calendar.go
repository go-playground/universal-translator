package lkt

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Wiótheȟika Wí", Feb: "Thiyóȟeyuŋka Wí", Mar: "Ištáwičhayazaŋ Wí", Apr: "Pȟežítȟo Wí", May: "Čhaŋwápetȟo Wí", Jun: "Wípazukȟa-wašté Wí", Jul: "Čhaŋpȟásapa Wí", Aug: "Wasútȟuŋ Wí", Sep: "Čhaŋwápeǧi Wí", Oct: "Čhaŋwápe-kasná Wí", Nov: "Waníyetu Wí", Dec: "Tȟahékapšuŋ Wí"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "A", Mon: "W", Tue: "N", Wed: "Y", Thu: "T", Fri: "Z", Sat: "O"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Aŋpétuwakȟaŋ", Mon: "Aŋpétuwaŋži", Tue: "Aŋpétunuŋpa", Wed: "Aŋpétuyamni", Thu: "Aŋpétutopa", Fri: "Aŋpétuzaptaŋ", Sat: "Owáŋgyužažapi"},
			},
			Periods: ut.CalendarPeriodFormatNames{},
		},
	}
}
