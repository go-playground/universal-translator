package ta

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "d/M/yy"},
			Time:     ut.CalendarDateFormat{Full: "a h:mm:ss zzzz", Long: "a h:mm:ss z", Medium: "a h:mm:ss", Short: "a h:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} ’அன்று’ {0}", Long: "{1} ’அன்று’ {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "ஜன.", Feb: "பிப்.", Mar: "மார்.", Apr: "ஏப்.", May: "மே", Jun: "ஜூன்", Jul: "ஜூலை", Aug: "ஆக.", Sep: "செப்.", Oct: "அக்.", Nov: "நவ.", Dec: "டிச."},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "ஜ", Feb: "பி", Mar: "மா", Apr: "ஏ", May: "மே", Jun: "ஜூ", Jul: "ஜூ", Aug: "ஆ", Sep: "செ", Oct: "அ", Nov: "ந", Dec: "டி"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "ஜனவரி", Feb: "பிப்ரவரி", Mar: "மார்ச்", Apr: "ஏப்ரல்", May: "மே", Jun: "ஜூன்", Jul: "ஜூலை", Aug: "ஆகஸ்டு", Sep: "செப்டம்பர்", Oct: "அக்டோபர்", Nov: "நவம்பர்", Dec: "டிசம்பர்"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "ஞாயி.", Mon: "திங்.", Tue: "செவ்.", Wed: "புத.", Thu: "வியா.", Fri: "வெள்.", Sat: "சனி"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "ஞா", Mon: "தி", Tue: "செ", Wed: "பு", Thu: "வி", Fri: "வெ", Sat: "ச"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "ஞா", Mon: "தி", Tue: "செ", Wed: "பு", Thu: "வி", Fri: "வெ", Sat: "ச"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "ஞாயிறு", Mon: "திங்கள்", Tue: "செவ்வாய்", Wed: "புதன்", Thu: "வியாழன்", Fri: "வெள்ளி", Sat: "சனி"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "முற்பகல்", PM: "பிற்பகல்"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "மு.ப", PM: "பி.ப"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "முற்பகல்", PM: "பிற்பகல்"},
			},
		},
	}
}
