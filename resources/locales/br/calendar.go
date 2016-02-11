package br

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{},
			Time:     ut.CalendarDateFormat{},
			DateTime: ut.CalendarDateFormat{Full: "{1} 'da' {0}", Long: "{1} 'da' {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Gen.", Feb: "Cʼhwe.", Mar: "Meur.", Apr: "Ebr.", May: "Mae", Jun: "Mezh.", Jul: "Goue.", Aug: "Eost", Sep: "Gwen.", Oct: "Here", Nov: "Du", Dec: "Ker."},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "01", Feb: "02", Mar: "03", Apr: "04", May: "05", Jun: "06", Jul: "07", Aug: "08", Sep: "09", Oct: "10", Nov: "11", Dec: "12"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Genver", Feb: "Cʼhwevrer", Mar: "Meurzh", Apr: "Ebrel", May: "Mae", Jun: "Mezheven", Jul: "Gouere", Aug: "Eost", Sep: "Gwengolo", Oct: "Here", Nov: "Du", Dec: "Kerzu"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Sul", Mon: "Lun", Tue: "Meu.", Wed: "Mer.", Thu: "Yaou", Fri: "Gwe.", Sat: "Sad."},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "Su", Mon: "L", Tue: "Mz", Wed: "Mc", Thu: "Y", Fri: "G", Sat: "Sa"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Sul", Mon: "Lun", Tue: "Meurzh", Wed: "Mercʼher", Thu: "Yaou", Fri: "Gwener", Sat: "Sadorn"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "A.M.", PM: "G.M."},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "am", PM: "gm"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "a-raok merenn", PM: "goude merenn"},
			},
		},
	}
}
