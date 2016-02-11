package chr

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
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "ᎤᏃ", Feb: "ᎧᎦ", Mar: "ᎠᏅ", Apr: "ᎧᏬ", May: "ᎠᏂ", Jun: "ᏕᎭ", Jul: "ᎫᏰ", Aug: "ᎦᎶ", Sep: "ᏚᎵ", Oct: "ᏚᏂ", Nov: "ᏅᏓ", Dec: "ᎥᏍ"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "Ꭴ", Feb: "Ꭷ", Mar: "Ꭰ", Apr: "Ꭷ", May: "Ꭰ", Jun: "Ꮥ", Jul: "Ꭻ", Aug: "Ꭶ", Sep: "Ꮪ", Oct: "Ꮪ", Nov: "Ꮕ", Dec: "Ꭵ"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "ᎤᏃᎸᏔᏅ", Feb: "ᎧᎦᎵ", Mar: "ᎠᏅᏱ", Apr: "ᎧᏬᏂ", May: "ᎠᏂᏍᎬᏘ", Jun: "ᏕᎭᎷᏱ", Jul: "ᎫᏰᏉᏂ", Aug: "ᎦᎶᏂ", Sep: "ᏚᎵᏍᏗ", Oct: "ᏚᏂᏅᏗ", Nov: "ᏅᏓᏕᏆ", Dec: "ᎥᏍᎩᏱ"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "ᏆᏍᎬ", Mon: "ᏉᏅᎯ", Tue: "ᏔᎵᏁ", Wed: "ᏦᎢᏁ", Thu: "ᏅᎩᏁ", Fri: "ᏧᎾᎩ", Sat: "ᏈᏕᎾ"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "Ꮖ", Mon: "Ꮙ", Tue: "Ꮤ", Wed: "Ꮶ", Thu: "Ꮕ", Fri: "Ꮷ", Sat: "Ꭴ"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "ᎤᎾᏙᏓᏆᏍᎬ", Mon: "ᎤᎾᏙᏓᏉᏅᎯ", Tue: "ᏔᎵᏁᎢᎦ", Wed: "ᏦᎢᏁᎢᎦ", Thu: "ᏅᎩᏁᎢᎦ", Fri: "ᏧᎾᎩᎶᏍᏗ", Sat: "ᎤᎾᏙᏓᏈᏕᎾ"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "ᏌᎾᎴ", PM: "ᏒᎯᏱᎢᏗᏢ"},
			},
		},
	}
}
