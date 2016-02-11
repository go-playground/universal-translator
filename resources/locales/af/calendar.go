package af

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, dd MMMM y", Long: "dd MMMM y", Medium: "dd MMM y", Short: "y-MM-dd"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Jan.", Feb: "Feb.", Mar: "Mrt.", Apr: "Apr.", May: "Mei", Jun: "Jun.", Jul: "Jul.", Aug: "Aug.", Sep: "Sep.", Oct: "Okt.", Nov: "Nov.", Dec: "Des."},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Januarie", Feb: "Februarie", Mar: "Maart", Apr: "April", May: "Mei", Jun: "Junie", Jul: "Julie", Aug: "Augustus", Sep: "September", Oct: "Oktober", Nov: "November", Dec: "Desember"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "So.", Mon: "Ma.", Tue: "Di.", Wed: "Wo.", Thu: "Do.", Fri: "Vr.", Sat: "Sa."},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "D", Wed: "W", Thu: "D", Fri: "V", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "So.", Mon: "Ma.", Tue: "Di.", Wed: "Wo.", Thu: "Do.", Fri: "Vr.", Sat: "Sa."},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Sondag", Mon: "Maandag", Tue: "Dinsdag", Wed: "Woensdag", Thu: "Donderdag", Fri: "Vrydag", Sat: "Saterdag"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "vm.", PM: "nm."},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "v", PM: "n"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "vm.", PM: "nm."},
			},
		},
	}
}
