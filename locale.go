package ut

// Locale contains all of the locale info needed for translation,
// number and date formatting
type Locale struct {
	Locale     string
	Number     Number
	Calendar   Calendar
	PluralRule string
}

// Number contains all of the number related formatting information
type Number struct {
	Symbols    Symbols
	Formats    NumberFormats
	Currencies []Currency
}

// Symbols contains the number symbols related to formatting
type Symbols struct {
	Decimal  string
	Group    string
	Negative string
	Percent  string
	PerMille string
}

// NumberFormats contains the number format information related to formatting
type NumberFormats struct {
	Decimal  string
	Currency string
	Percent  string
}

// Currency contains the Currency related formatting information
type Currency struct {
	Currency    string
	DisplayName string
	Symbol      string
}

// Calendar contains the DateTime formatting information
type Calendar struct {
	Formats     CalendarFormats
	FormatNames CalendarFormatNames
}

// CalendarFormats contains the DateTime format information
type CalendarFormats struct {
	Date     CalendarDateFormat
	Time     CalendarDateFormat
	DateTime CalendarDateFormat
}

// CalendarDateFormat contains the DateTime length format information
type CalendarDateFormat struct{ Full, Long, Medium, Short string }

// CalendarFormatNames contains the DateTime name information
type CalendarFormatNames struct {
	Months  CalendarMonthFormatNames
	Days    CalendarDayFormatNames
	Periods CalendarPeriodFormatNames
}

// CalendarMonthFormatNames contains the DateTime month formats information
type CalendarMonthFormatNames struct {
	Abbreviated CalendarMonthFormatNameValue
	Narrow      CalendarMonthFormatNameValue
	Short       CalendarMonthFormatNameValue
	Wide        CalendarMonthFormatNameValue
}

// CalendarMonthFormatNameValue contains the DateTime month name information
type CalendarMonthFormatNameValue struct {
	Jan, Feb, Mar, Apr, May, Jun, Jul, Aug, Sep, Oct, Nov, Dec string
}

// CalendarDayFormatNames contains the DateTime month name information
type CalendarDayFormatNames struct {
	Abbreviated CalendarDayFormatNameValue
	Narrow      CalendarDayFormatNameValue
	Short       CalendarDayFormatNameValue
	Wide        CalendarDayFormatNameValue
}

// CalendarDayFormatNameValue contains the DateTime day name information
type CalendarDayFormatNameValue struct {
	Sun, Mon, Tue, Wed, Thu, Fri, Sat string
}

// CalendarPeriodFormatNames contains the DateTime period information
type CalendarPeriodFormatNames struct {
	Abbreviated CalendarPeriodFormatNameValue
	Narrow      CalendarPeriodFormatNameValue
	Short       CalendarPeriodFormatNameValue
	Wide        CalendarPeriodFormatNameValue
}

// CalendarPeriodFormatNameValue contains the DateTime period name information
type CalendarPeriodFormatNameValue struct {
	AM, PM string
}
