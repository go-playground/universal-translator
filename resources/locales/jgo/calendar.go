package jgo

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, y MMMM dd", Long: "y MMMM d", Medium: "y MMM d", Short: "y-MM-dd"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Nduŋmbi Saŋ", Feb: "Pɛsaŋ Pɛ́pá", Mar: "Pɛsaŋ Pɛ́tát", Apr: "Pɛsaŋ Pɛ́nɛ́kwa", May: "Pɛsaŋ Pataa", Jun: "Pɛsaŋ Pɛ́nɛ́ntúkú", Jul: "Pɛsaŋ Saambá", Aug: "Pɛsaŋ Pɛ́nɛ́fɔm", Sep: "Pɛsaŋ Pɛ́nɛ́pfúꞋú", Oct: "Pɛsaŋ Nɛgɛ́m", Nov: "Pɛsaŋ Ntsɔ̌pmɔ́", Dec: "Pɛsaŋ Ntsɔ̌ppá"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Nduŋmbi Saŋ", Feb: "Pɛsaŋ Pɛ́pá", Mar: "Pɛsaŋ Pɛ́tát", Apr: "Pɛsaŋ Pɛ́nɛ́kwa", May: "Pɛsaŋ Pataa", Jun: "Pɛsaŋ Pɛ́nɛ́ntúkú", Jul: "Pɛsaŋ Saambá", Aug: "Pɛsaŋ Pɛ́nɛ́fɔm", Sep: "Pɛsaŋ Pɛ́nɛ́pfúꞋú", Oct: "Pɛsaŋ Nɛgɛ́m", Nov: "Pɛsaŋ Ntsɔ̌pmɔ́", Dec: "Pɛsaŋ Ntsɔ̌ppá"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "Sɔ́ndi", Mon: "Mɔ́ndi", Tue: "Ápta Mɔ́ndi", Wed: "Wɛ́nɛsɛdɛ", Thu: "Tɔ́sɛdɛ", Fri: "Fɛlâyɛdɛ", Sat: "Sásidɛ"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "Sɔ́", Mon: "Mɔ́", Tue: "ÁM", Wed: "Wɛ́", Thu: "Tɔ́", Fri: "Fɛ", Sat: "Sá"},
				Short:       ut.CalendarDayFormatNameValue{},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Sɔ́ndi", Mon: "Mɔ́ndi", Tue: "Ápta Mɔ́ndi", Wed: "Wɛ́nɛsɛdɛ", Thu: "Tɔ́sɛdɛ", Fri: "Fɛlâyɛdɛ", Sat: "Sásidɛ"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "mbaꞌmbaꞌ", PM: "ŋka mbɔ́t nji"},
			},
		},
	}
}
