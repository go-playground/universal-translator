package or

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d-M-yy"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "ଜାନୁଆରୀ", Feb: "ଫେବୃଆରୀ", Mar: "ମାର୍ଚ୍ଚ", Apr: "ଅପ୍ରେଲ", May: "ମଇ", Jun: "ଜୁନ", Jul: "ଜୁଲାଇ", Aug: "ଅଗଷ୍ଟ", Sep: "ସେପ୍ଟେମ୍ବର", Oct: "ଅକ୍ଟୋବର", Nov: "ନଭେମ୍ବର", Dec: "ଡିସେମ୍ବର"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "ଜା", Feb: "ଫେ", Mar: "ମା", Apr: "ଅ", May: "ମଇ", Jun: "ଜୁ", Jul: "ଜୁ", Aug: "ଅ", Sep: "ସେ", Oct: "ଅ", Nov: "ନ", Dec: "ଡି"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "ଜାନୁଆରୀ", Feb: "ଫେବୃଆରୀ", Mar: "ମାର୍ଚ୍ଚ", Apr: "ଅପ୍ରେଲ", May: "ମଇ", Jun: "ଜୁନ", Jul: "ଜୁଲାଇ", Aug: "ଅଗଷ୍ଟ", Sep: "ସେପ୍ଟେମ୍ବର", Oct: "ଅକ୍ଟୋବର", Nov: "ନଭେମ୍ବର", Dec: "ଡିସେମ୍ବର"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "ରବି", Mon: "ସୋମ", Tue: "ମଙ୍ଗଳ", Wed: "ବୁଧ", Thu: "ଗୁରୁ", Fri: "ଶୁକ୍ର", Sat: "ଶନି"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "ର", Mon: "ସୋ", Tue: "ମ", Wed: "ବୁ", Thu: "ଗୁ", Fri: "ଶୁ", Sat: "ଶ"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "ରବିବାର", Mon: "ସୋମବାର", Tue: "ମଙ୍ଗଳବାର", Wed: "ବୁଧବାର", Thu: "ଗୁରୁବାର", Fri: "ଶୁକ୍ରବାର", Sat: "ଶନିବାର"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "am", PM: "pm"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "am", PM: "pm"},
			},
		},
	}
}
