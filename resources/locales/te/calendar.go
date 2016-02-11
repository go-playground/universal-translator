package te

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "d, MMMM y, EEEE", Long: "d MMMM, y", Medium: "d MMM, y", Short: "dd-MM-yy"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "జన", Feb: "ఫిబ్ర", Mar: "మార్చి", Apr: "ఏప్రి", May: "మే", Jun: "జూన్", Jul: "జులై", Aug: "ఆగస్టు", Sep: "సెప్టెం", Oct: "అక్టో", Nov: "నవం", Dec: "డిసెం"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "జ", Feb: "ఫి", Mar: "మా", Apr: "ఏ", May: "మే", Jun: "జూ", Jul: "జు", Aug: "ఆ", Sep: "సె", Oct: "అ", Nov: "న", Dec: "డి"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "జనవరి", Feb: "ఫిబ్రవరి", Mar: "మార్చి", Apr: "ఏప్రిల్", May: "మే", Jun: "జూన్", Jul: "జులై", Aug: "ఆగస్టు", Sep: "సెప్టెంబర్", Oct: "అక్టోబర్", Nov: "నవంబర్", Dec: "డిసెంబర్"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "ఆది", Mon: "సోమ", Tue: "మంగళ", Wed: "బుధ", Thu: "గురు", Fri: "శుక్ర", Sat: "శని"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "ఆ", Mon: "సో", Tue: "మ", Wed: "బు", Thu: "గు", Fri: "శు", Sat: "శ"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "ఆది", Mon: "సోమ", Tue: "మం", Wed: "బుధ", Thu: "గురు", Fri: "శుక్ర", Sat: "శని"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "ఆదివారం", Mon: "సోమవారం", Tue: "మంగళవారం", Wed: "బుధవారం", Thu: "గురువారం", Fri: "శుక్రవారం", Sat: "శనివారం"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
			},
		},
	}
}
