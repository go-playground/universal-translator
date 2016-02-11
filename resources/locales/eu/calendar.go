package eu

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "y('e')'ko' MMMM d, EEEE", Long: "y('e')'ko' MMMM d", Medium: "y MMM d", Short: "y/MM/dd"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss (zzzz)", Long: "HH:mm:ss (z)", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Urt.", Feb: "Ots.", Mar: "Mar.", Apr: "Api.", May: "Mai.", Jun: "Eka.", Jul: "Uzt.", Aug: "Abu.", Sep: "Ira.", Oct: "Urr.", Nov: "Aza.", Dec: "Abe."},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "U", Feb: "O", Mar: "M", Apr: "A", May: "M", Jun: "E", Jul: "U", Aug: "A", Sep: "I", Oct: "U", Nov: "A", Dec: "A"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Urtarrila", Feb: "Otsaila", Mar: "Martxoa", Apr: "Apirila", May: "Maiatza", Jun: "Ekaina", Jul: "Uztaila", Aug: "Abuztua", Sep: "Iraila", Oct: "Urria", Nov: "Azaroa", Dec: "Abendua"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Ig.", Mon: "Al.", Tue: "Ar.", Wed: "Az.", Thu: "Og.", Fri: "Or.", Sat: "Lr."},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "I", Mon: "A", Tue: "A", Wed: "A", Thu: "O", Fri: "O", Sat: "L"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "ig.", Mon: "al.", Tue: "ar.", Wed: "az.", Thu: "og.", Fri: "or.", Sat: "lr."},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Igandea", Mon: "Astelehena", Tue: "Asteartea", Wed: "Asteazkena", Thu: "Osteguna", Fri: "Ostirala", Sat: "Larunbata"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "g", PM: "a"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			},
		},
	}
}
