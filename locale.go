package ut

import "time"

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
	Currencies CurrencyFormatValue
}

// CurrencyFormatValue contains the currency information
type CurrencyFormatValue map[string]Currency

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
	Decimal            string
	Currency           string
	CurrencyAccounting string
	Percent            string
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
	DateEra  DateEra
	Time     CalendarDateFormat
	DateTime CalendarDateFormat
}

// DateEra contains Era Date Formats as they can vary if AD vs BC
type DateEra struct {
	BC CalendarDateFormat
	AD CalendarDateFormat
}

// CalendarDateFormat contains the DateTime length format information
type CalendarDateFormat struct{ Full, Long, Medium, Short string }

// CalendarFormatNames contains the DateTime name information
type CalendarFormatNames struct {
	Months  CalendarMonthFormatNames
	Days    CalendarDayFormatNames
	Periods CalendarPeriodFormatNames
	Eras    Eras
}

// Eras contains the Era information for both AD and BC
type Eras struct {
	AD CalendarEraFormatNames
	BC CalendarEraFormatNames
}

// CalendarEraFormatNames contains the Era name information
type CalendarEraFormatNames struct{ Full, Abbrev, Narrow string }

// CalendarMonthFormatNames contains the DateTime month formats information
type CalendarMonthFormatNames struct {
	Abbreviated CalendarMonthFormatNameValue
	Narrow      CalendarMonthFormatNameValue
	Short       CalendarMonthFormatNameValue
	Wide        CalendarMonthFormatNameValue
}

// CalendarMonthFormatNameValue contains the DateTime month name information
type CalendarMonthFormatNameValue map[time.Month]string

// CalendarDayFormatNames contains the DateTime month name information
type CalendarDayFormatNames struct {
	Abbreviated CalendarDayFormatNameValue
	Narrow      CalendarDayFormatNameValue
	Short       CalendarDayFormatNameValue
	Wide        CalendarDayFormatNameValue
}

// CalendarDayFormatNameValue contains the DateTime day name information
type CalendarDayFormatNameValue map[time.Weekday]string

// CalendarPeriodFormatNames contains the DateTime period information
type CalendarPeriodFormatNames struct {
	Abbreviated CalendarPeriodFormatNameValue
	Narrow      CalendarPeriodFormatNameValue
	Short       CalendarPeriodFormatNameValue
	Wide        CalendarPeriodFormatNameValue
}

// CalendarPeriodFormatNameValue contains the DateTime period name information
type CalendarPeriodFormatNameValue map[string]string
