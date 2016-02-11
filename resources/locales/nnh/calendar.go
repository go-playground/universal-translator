package nnh

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE , 'lyɛ'̌ʼ d 'na' MMMM, y", Long: "'lyɛ'̌ʼ d 'na' MMMM, y", Medium: "d MMM, y", Short: "dd/MM/yy"},
			Time:     ut.CalendarDateFormat{},
			DateTime: ut.CalendarDateFormat{Full: "{1},{0}", Long: "{1}, {0}", Medium: "", Short: ""},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "saŋ tsetsɛ̀ɛ lùm", Feb: "saŋ kàg ngwóŋ", Mar: "saŋ lepyè shúm", Apr: "saŋ cÿó", May: "saŋ tsɛ̀ɛ cÿó", Jun: "saŋ njÿoláʼ", Jul: "saŋ tyɛ̀b tyɛ̀b mbʉ̀ŋ", Aug: "saŋ mbʉ̀ŋ", Sep: "saŋ ngwɔ̀ʼ mbÿɛ", Oct: "saŋ tàŋa tsetsáʼ", Nov: "saŋ mejwoŋó", Dec: "saŋ lùm"},
				Narrow:      ut.CalendarMonthFormatNameValue{},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "saŋ tsetsɛ̀ɛ lùm", Feb: "saŋ kàg ngwóŋ", Mar: "saŋ lepyè shúm", Apr: "saŋ cÿó", May: "saŋ tsɛ̀ɛ cÿó", Jun: "saŋ njÿoláʼ", Jul: "saŋ tyɛ̀b tyɛ̀b mbʉ̀ŋ", Aug: "saŋ mbʉ̀ŋ", Sep: "saŋ ngwɔ̀ʼ mbÿɛ", Oct: "saŋ tàŋa tsetsáʼ", Nov: "saŋ mejwoŋó", Dec: "saŋ lùm"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "lyɛʼɛ́ sẅíŋtè", Mon: "mvfò lyɛ̌ʼ", Tue: "mbɔ́ɔntè mvfò lyɛ̌ʼ", Wed: "tsètsɛ̀ɛ lyɛ̌ʼ", Thu: "mbɔ́ɔntè tsetsɛ̀ɛ lyɛ̌ʼ", Fri: "mvfò màga lyɛ̌ʼ", Sat: "màga lyɛ̌ʼ"},
				Narrow:      ut.CalendarDayFormatNameValue{},
				Short:       ut.CalendarDayFormatNameValue{Sun: "lyɛʼɛ́ sẅíŋtè", Mon: "mvfò lyɛ̌ʼ", Tue: "mbɔ́ɔntè mvfò lyɛ̌ʼ", Wed: "tsètsɛ̀ɛ lyɛ̌ʼ", Thu: "mbɔ́ɔntè tsetsɛ̀ɛ lyɛ̌ʼ", Fri: "mvfò màga lyɛ̌ʼ", Sat: "màga lyɛ̌ʼ"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "lyɛʼɛ́ sẅíŋtè", Mon: "mvfò lyɛ̌ʼ", Tue: "mbɔ́ɔntè mvfò lyɛ̌ʼ", Wed: "tsètsɛ̀ɛ lyɛ̌ʼ", Thu: "mbɔ́ɔntè tsetsɛ̀ɛ lyɛ̌ʼ", Fri: "mvfò màga lyɛ̌ʼ", Sat: "màga lyɛ̌ʼ"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "mbaʼámbaʼ", PM: "ncwònzém"},
			},
		},
	}
}
