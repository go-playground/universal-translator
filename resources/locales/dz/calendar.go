package dz

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, སྤྱི་ལོ་y MMMM ཚེས་dd", Long: "སྤྱི་ལོ་y MMMM ཚེས་ dd", Medium: "སྤྱི་ལོ་y ཟླ་MMM ཚེས་dd", Short: "y-MM-dd"},
			Time:     ut.CalendarDateFormat{Full: "ཆུ་ཚོད་ h སྐར་མ་ mm:ss a zzzz", Long: "ཆུ་ཚོད་ h སྐར་མ་ mm:ss a z", Medium: "ཆུ་ཚོད་h:mm:ss a", Short: "ཆུ་ཚོད་ h སྐར་མ་ mm a"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "ཟླ་༡", Feb: "ཟླ་༢", Mar: "ཟླ་༣", Apr: "ཟླ་༤", May: "ཟླ་༥", Jun: "ཟླ་༦", Jul: "ཟླ་༧", Aug: "ཟླ་༨", Sep: "ཟླ་༩", Oct: "ཟླ་༡༠", Nov: "ཟླ་༡༡", Dec: "ཟླ་༡༢"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "༡", Feb: "༢", Mar: "༣", Apr: "༤", May: "༥", Jun: "༦", Jul: "༧", Aug: "༨", Sep: "༩", Oct: "༡༠", Nov: "༡༡", Dec: "༡༢"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "སྤྱི་ཟླ་དངཔ་", Feb: "སྤྱི་ཟླ་གཉིས་པ་", Mar: "སྤྱི་ཟླ་གསུམ་པ་", Apr: "སྤྱི་ཟླ་བཞི་པ", May: "སྤྱི་ཟླ་ལྔ་པ་", Jun: "སྤྱི་ཟླ་དྲུག་པ", Jul: "སྤྱི་ཟླ་བདུན་པ་", Aug: "སྤྱི་ཟླ་བརྒྱད་པ་", Sep: "སྤྱི་ཟླ་དགུ་པ་", Oct: "སྤྱི་ཟླ་བཅུ་པ་", Nov: "སྤྱི་ཟླ་བཅུ་གཅིག་པ་", Dec: "སྤྱི་ཟླ་བཅུ་གཉིས་པ་"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "ཟླ་", Mon: "མིར་", Tue: "ལྷག་", Wed: "ཕུར་", Thu: "སངས་", Fri: "སྤེན་", Sat: "ཉི་"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "ཟླ", Mon: "མིར", Tue: "ལྷག", Wed: "ཕུར", Thu: "སངྶ", Fri: "སྤེན", Sat: "ཉི"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "ཟླ་", Mon: "མིར་", Tue: "ལྷག་", Wed: "ཕུར་", Thu: "སངས་", Fri: "སྤེན་", Sat: "ཉི་"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "གཟའ་ཟླ་བ་", Mon: "གཟའ་མིག་དམར་", Tue: "གཟའ་ལྷག་པ་", Wed: "གཟའ་ཕུར་བུ་", Thu: "གཟའ་པ་སངས་", Fri: "གཟའ་སྤེན་པ་", Sat: "གཟའ་ཉི་མ་"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "སྔ་ཆ་", PM: "ཕྱི་ཆ་"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "སྔ་ཆ་", PM: "ཕྱི་ཆ་"},
			},
		},
	}
}
