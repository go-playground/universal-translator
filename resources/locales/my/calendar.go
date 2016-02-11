package my

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE၊ dd MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd-MM-yy"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1}မှာ {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "ဇန်", Feb: "ဖေ", Mar: "မတ်", Apr: "ဧပြီ", May: "မေ", Jun: "ဇွန်", Jul: "ဇူ", Aug: "ဩ", Sep: "စက်", Oct: "အောက်", Nov: "နို", Dec: "ဒီ"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "ဇ", Feb: "ဖ", Mar: "မ", Apr: "ဧ", May: "မ", Jun: "ဇ", Jul: "ဇ", Aug: "ဩ", Sep: "စ", Oct: "အ", Nov: "န", Dec: "ဒ"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "ဇန်နဝါရီ", Feb: "ဖေဖော်ဝါရီ", Mar: "မတ်", Apr: "ဧပြီ", May: "မေ", Jun: "ဇွန်", Jul: "ဇူလိုင်", Aug: "ဩဂုတ်", Sep: "စက်တင်ဘာ", Oct: "အောက်တိုဘာ", Nov: "နိုဝင်ဘာ", Dec: "ဒီဇင်ဘာ"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "တနင်္ဂနွေ", Mon: "တနင်္လာ", Tue: "အင်္ဂါ", Wed: "ဗုဒ္ဓဟူး", Thu: "ကြာသပတေး", Fri: "သောကြာ", Sat: "စနေ"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "တ", Mon: "တ", Tue: "အ", Wed: "ဗ", Thu: "က", Fri: "သ", Sat: "စ"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "တနင်္ဂနွေ", Mon: "တနင်္လာ", Tue: "အင်္ဂါ", Wed: "ဗုဒ္ဓဟူး", Thu: "ကြာသပတေး", Fri: "သောကြာ", Sat: "စနေ"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "တနင်္ဂနွေ", Mon: "တနင်္လာ", Tue: "အင်္ဂါ", Wed: "ဗုဒ္ဓဟူး", Thu: "ကြာသပတေး", Fri: "သောကြာ", Sat: "စနေ"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "နံနက်", PM: "ညနေ"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "နံနက်", PM: "ညနေ"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "နံနက်", PM: "ညနေ"},
			},
		},
	}
}
