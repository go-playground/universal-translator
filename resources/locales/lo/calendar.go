package lo

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE ທີ d MMMM G y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"},
			Time:     ut.CalendarDateFormat{Full: "H ໂມງ m ນາທີ ss ວິນາທີ zzzz", Long: "H ໂມງ m ນາທີ ss ວິນາທີ z", Medium: "H:mm:ss", Short: "H:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "ມ.ກ.", Feb: "ກ.ພ.", Mar: "ມ.ນ.", Apr: "ມ.ສ.", May: "ພ.ພ.", Jun: "ມິ.ຖ.", Jul: "ກ.ລ.", Aug: "ສ.ຫ.", Sep: "ກ.ຍ.", Oct: "ຕ.ລ.", Nov: "ພ.ຈ.", Dec: "ທ.ວ."},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "ມັງກອນ", Feb: "ກຸມພາ", Mar: "ມີນາ", Apr: "ເມສາ", May: "ພຶດສະພາ", Jun: "ມິຖຸນາ", Jul: "ກໍລະກົດ", Aug: "ສິງຫາ", Sep: "ກັນຍາ", Oct: "ຕຸລາ", Nov: "ພະຈິກ", Dec: "ທັນວາ"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "ອາທິດ", Mon: "ຈັນ", Tue: "ອັງຄານ", Wed: "ພຸດ", Thu: "ພະຫັດ", Fri: "ສຸກ", Sat: "ເສົາ"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "ອ", Mon: "ຈ", Tue: "ອ", Wed: "ພ", Thu: "ພຫ", Fri: "ສຸ", Sat: "ສ"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "ອາ.", Mon: "ຈ.", Tue: "ອ.", Wed: "ພ.", Thu: "ພຫ.", Fri: "ສຸ.", Sat: "ສ."},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "ວັນອາທິດ", Mon: "ວັນຈັນ", Tue: "ວັນອັງຄານ", Wed: "ວັນພຸດ", Thu: "ວັນພະຫັດ", Fri: "ວັນສຸກ", Sat: "ວັນເສົາ"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "ກ່ອນທ່ຽງ", PM: "ຫຼັງທ່ຽງ"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "ກທ", PM: "ຫຼທ"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "ກ່ອນທ່ຽງ", PM: "ຫຼັງທ່ຽງ"},
			},
		},
	}
}
