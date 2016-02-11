package ka

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, dd MMMM, y", Long: "d MMMM, y", Medium: "d MMM. y", Short: "dd.MM.yy"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "იან", Feb: "თებ", Mar: "მარ", Apr: "აპრ", May: "მაი", Jun: "ივნ", Jul: "ივლ", Aug: "აგვ", Sep: "სექ", Oct: "ოქტ", Nov: "ნოე", Dec: "დეკ"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "ი", Feb: "თ", Mar: "მ", Apr: "ა", May: "მ", Jun: "ი", Jul: "ი", Aug: "ა", Sep: "ს", Oct: "ო", Nov: "ნ", Dec: "დ"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "იანვარი", Feb: "თებერვალი", Mar: "მარტი", Apr: "აპრილი", May: "მაისი", Jun: "ივნისი", Jul: "ივლისი", Aug: "აგვისტო", Sep: "სექტემბერი", Oct: "ოქტომბერი", Nov: "ნოემბერი", Dec: "დეკემბერი"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "კვი", Mon: "ორშ", Tue: "სამ", Wed: "ოთხ", Thu: "ხუთ", Fri: "პარ", Sat: "შაბ"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "კ", Mon: "ო", Tue: "ს", Wed: "ო", Thu: "ხ", Fri: "პ", Sat: "შ"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "კვ", Mon: "ორ", Tue: "სმ", Wed: "ოთ", Thu: "ხთ", Fri: "პრ", Sat: "შბ"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "კვირა", Mon: "ორშაბათი", Tue: "სამშაბათი", Wed: "ოთხშაბათი", Thu: "ხუთშაბათი", Fri: "პარასკევი", Sat: "შაბათი"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "a", PM: "p"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "შუადღემდე", PM: "შუადღ. შემდეგ"},
			},
		},
	}
}
