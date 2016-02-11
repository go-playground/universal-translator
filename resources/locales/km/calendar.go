package km

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/yy"},
			Time:     ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"},
			DateTime: ut.CalendarDateFormat{Full: "{1} នៅ {0}", Long: "{1} នៅ {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "មករា", Feb: "កុម្ភៈ", Mar: "មីនា", Apr: "មេសា", May: "ឧសភា", Jun: "មិថុនា", Jul: "កក្កដា", Aug: "សីហា", Sep: "កញ្ញា", Oct: "តុលា", Nov: "វិច្ឆិកា", Dec: "ធ្នូ"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "មករា", Feb: "កុម្ភៈ", Mar: "មីនា", Apr: "មេសា", May: "ឧសភា", Jun: "មិថុនា", Jul: "កក្កដា", Aug: "សីហា", Sep: "កញ្ញា", Oct: "តុលា", Nov: "វិច្ឆិកា", Dec: "ធ្នូ"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "អាទិត្យ", Mon: "ច័ន្ទ", Tue: "អង្គារ", Wed: "ពុធ", Thu: "ព្រហស្បតិ៍", Fri: "សុក្រ", Sat: "សៅរ៍"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "អា", Mon: "ច", Tue: "អ", Wed: "ពុ", Thu: "ព្រ", Fri: "សុ", Sat: "ស"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "អាទិត្យ", Mon: "ច័ន្ទ", Tue: "អង្គារ", Wed: "ពុធ", Thu: "ព្រហស្បតិ៍", Fri: "សុក្រ", Sat: "សៅរ៍"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "អាទិត្យ", Mon: "ច័ន្ទ", Tue: "អង្គារ", Wed: "ពុធ", Thu: "ព្រហស្បតិ៍", Fri: "សុក្រ", Sat: "សៅរ៍"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "ព្រឹក", PM: "ល្ងាច"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "ព្រឹក", PM: "ល្ងាច"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "ព្រឹក", PM: "ល្ងាច"},
			},
		},
	}
}
