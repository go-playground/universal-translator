package bo

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "y MMMMའི་ཚེས་d, EEEE", Long: "སྤྱི་ལོ་y MMMMའི་ཚེས་d", Medium: "y ལོའི་MMMཚེས་d", Short: "y-MM-dd"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "ཟླ་༡", Feb: "ཟླ་༢", Mar: "ཟླ་༣", Apr: "ཟླ་༤", May: "ཟླ་༥", Jun: "ཟླ་༦", Jul: "ཟླ་༧", Aug: "ཟླ་༨", Sep: "ཟླ་༩", Oct: "ཟླ་༡༠", Nov: "ཟླ་༡༡", Dec: "ཟླ་༡༢"},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "ཟླ་བ་དང་པོ་", Feb: "ཟླ་བ་གཉིས་པ་", Mar: "ཟླ་བ་གསུམ་པ་", Apr: "ཟླ་བ་བཞི་པ་", May: "ཟླ་བ་ལྔ་པ་", Jun: "ཟླ་བ་དྲུག་པ་", Jul: "ཟླ་བ་བདུན་པ་", Aug: "ཟླ་བ་བརྒྱད་པ་", Sep: "ཟླ་བ་དགུ་པ་", Oct: "ཟླ་བ་བཅུ་པ་", Nov: "ཟླ་བ་བཅུ་གཅིག་པ་", Dec: "ཟླ་བ་བཅུ་གཉིས་པ་"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "ཉི་མ་", Mon: "ཟླ་བ་", Tue: "མིག་དམར་", Wed: "ལྷག་པ་", Thu: "ཕུར་བུ་", Fri: "པ་སངས་", Sat: "སྤེན་པ་"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "ཉི", Mon: "ཟླ", Tue: "མིག", Wed: "ལྷག", Thu: "ཕུར", Fri: "སངས", Sat: "སྤེན"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "གཟའ་ཉི་མ་", Mon: "གཟའ་ཟླ་བ་", Tue: "གཟའ་མིག་དམར་", Wed: "གཟའ་ལྷག་པ་", Thu: "གཟའ་ཕུར་བུ་", Fri: "གཟའ་པ་སངས་", Sat: "གཟའ་སྤེན་པ་"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "སྔ་དྲོ་", PM: "ཕྱི་དྲོ་"},
			},
		},
	}
}
