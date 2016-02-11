package ksh

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, 'dä' d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM. y", Short: "d. M. y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Jan.", Feb: "Fäb.", Mar: "Mäz.", Apr: "Apr.", May: "Mäi", Jun: "Jun.", Jul: "Jul.", Aug: "Ouj.", Sep: "Säp.", Oct: "Okt.", Nov: "Nov.", Dec: "Dez."},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Jannewa", Feb: "Fäbrowa", Mar: "Määz", Apr: "Aprell", May: "Mäi", Jun: "Juuni", Jul: "Juuli", Aug: "Oujoß", Sep: "Septämber", Oct: "Oktoober", Nov: "Novämber", Dec: "Dezämber"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Su.", Mon: "Mo.", Tue: "Di.", Wed: "Me.", Thu: "Du.", Fri: "Fr.", Sat: "Sa."},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "D", Wed: "M", Thu: "D", Fri: "F", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "Su", Mon: "Mo", Tue: "Di", Wed: "Me", Thu: "Du", Fri: "Fr", Sat: "Sa"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Sunndaach", Mon: "Moondaach", Tue: "Dinnsdaach", Wed: "Metwoch", Thu: "Dunnersdaach", Fri: "Friidaach", Sat: "Samsdaach"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "v.m.", PM: "n.m."},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "Vormittag", PM: "Nachmittag"},
			},
		},
	}
}
