package nmg

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "ng1", Feb: "ng2", Mar: "ng3", Apr: "ng4", May: "ng5", Jun: "ng6", Jul: "ng7", Aug: "ng8", Sep: "ng9", Oct: "ng10", Nov: "ng11", Dec: "kris"},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "ngwɛn matáhra", Feb: "ngwɛn ńmba", Mar: "ngwɛn ńlal", Apr: "ngwɛn ńna", May: "ngwɛn ńtan", Jun: "ngwɛn ńtuó", Jul: "ngwɛn hɛmbuɛrí", Aug: "ngwɛn lɔmbi", Sep: "ngwɛn rɛbvuâ", Oct: "ngwɛn wum", Nov: "ngwɛn wum navǔr", Dec: "krísimin"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "sɔ́n", Mon: "mɔ́n", Tue: "smb", Wed: "sml", Thu: "smn", Fri: "mbs", Sat: "sas"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "s", Mon: "m", Tue: "s", Wed: "s", Thu: "s", Fri: "m", Sat: "s"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "sɔ́ndɔ", Mon: "mɔ́ndɔ", Tue: "sɔ́ndɔ mafú mába", Wed: "sɔ́ndɔ mafú málal", Thu: "sɔ́ndɔ mafú mána", Fri: "mabágá má sukul", Sat: "sásadi"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "maná", PM: "kugú"},
			},
		},
	}
}
