package ut_test

import (
	"testing"
	"time"

	"github.com/go-playground/universal-translator"
	_ "github.com/go-playground/universal-translator/resources/locales"

	. "gopkg.in/go-playground/assert.v1"
)

// NOTES:
// - Run "go test" to run tests
// - Run "gocov test | gocov report" to report on test converage by file
// - Run "gocov test | gocov annotate -" to report on all code and functions, those ,marked with "MISS" were never called
//
// or
//
// -- may be a good idea to change to output path to somewherelike /tmp
// go test -coverprofile cover.out && go tool cover -html=cover.out -o cover.html
//

func TestDateTimeEn(t *testing.T) {

	en, err := ut.GetTranslator("en")
	Equal(t, err, nil)

	datetime, err := time.Parse(dateTimeString, dateTimeString)
	Equal(t, err, nil)

	// test the public method
	dt, err := en.FmtDateFullSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "Monday, January 2, 2006")

	dt, err = en.FmtDateFullSafe(bc300DateTime)
	Equal(t, err, nil)
	Equal(t, dt, "Saturday, January 2, 300 Before Christ")

	dt = en.FmtDateFull(datetime)
	Equal(t, dt, "Monday, January 2, 2006")

	dt = en.FmtDateFull(bc300DateTime)
	Equal(t, dt, "Saturday, January 2, 300 Before Christ")

	dt, err = en.FmtDateLongSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "January 2, 2006")

	dt, err = en.FmtDateLongSafe(bc300DateTime)
	Equal(t, err, nil)
	Equal(t, dt, "January 2, 300 BC")

	dt = en.FmtDateLong(datetime)
	Equal(t, dt, "January 2, 2006")

	dt = en.FmtDateLong(bc300DateTime)
	Equal(t, dt, "January 2, 300 BC")

	dt, err = en.FmtDateMediumSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "Jan 2, 2006")

	dt = en.FmtDateMedium(datetime)
	Equal(t, dt, "Jan 2, 2006")

	dt, err = en.FmtDateShortSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "1/2/06")

	dt = en.FmtDateShort(datetime)
	Equal(t, dt, "1/2/06")

	dt, err = en.FmtTimeFullSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "3:04:05 PM")

	dt = en.FmtTimeFull(datetime)
	Equal(t, dt, "3:04:05 PM")

	dt, err = en.FmtTimeLongSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "3:04:05 PM")

	dt = en.FmtTimeLong(datetime)
	Equal(t, dt, "3:04:05 PM")

	dt, err = en.FmtTimeMediumSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "3:04:05 PM")

	dt = en.FmtTimeMedium(datetime)
	Equal(t, dt, "3:04:05 PM")

	dt, err = en.FmtTimeShortSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "3:04 PM")

	dt = en.FmtTimeShort(datetime)
	Equal(t, dt, "3:04 PM")

	dt, err = en.FmtDateTimeFullSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "Monday, January 2, 2006 at 3:04:05 PM")

	dt, err = en.FmtDateTimeFullSafe(bc300DateTime)
	Equal(t, err, nil)
	Equal(t, dt, "Saturday, January 2, 300 Before Christ at 3:04:05 AM")

	dt = en.FmtDateTimeFull(datetime)
	Equal(t, dt, "Monday, January 2, 2006 at 3:04:05 PM")

	dt = en.FmtDateTimeFull(bc300DateTime)
	Equal(t, dt, "Saturday, January 2, 300 Before Christ at 3:04:05 AM")

	dt, err = en.FmtDateTimeLongSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "January 2, 2006 at 3:04:05 PM")

	dt, err = en.FmtDateTimeLongSafe(bc300DateTime)
	Equal(t, err, nil)
	Equal(t, dt, "January 2, 300 BC at 3:04:05 AM")

	dt = en.FmtDateTimeLong(datetime)
	Equal(t, dt, "January 2, 2006 at 3:04:05 PM")

	dt = en.FmtDateTimeLong(bc300DateTime)
	Equal(t, dt, "January 2, 300 BC at 3:04:05 AM")

	dt, err = en.FmtDateTimeMediumSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "Jan 2, 2006, 3:04:05 PM")

	dt = en.FmtDateTimeMedium(datetime)
	Equal(t, dt, "Jan 2, 2006, 3:04:05 PM")

	dt, err = en.FmtDateTimeShortSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "1/2/06, 3:04 PM")

	dt = en.FmtDateTimeShort(datetime)
	Equal(t, dt, "1/2/06, 3:04 PM")

	dt, err = en.FmtDateTimeSafe(datetime, "MMMM d yy")
	Equal(t, err, nil)
	Equal(t, dt, "January 2 06")

	dt = en.FmtDateTime(datetime, "MMMM d yy")
	Equal(t, err, nil)
	Equal(t, dt, "January 2 06")

	dt, err = en.FmtDateTimeSafe(datetime, "not a date pattern")
	NotEqual(t, err, nil)
	Equal(t, dt, "")
	Equal(t, err.Error(), "unknown datetime format unit: n")
}

func TestCurrencyEn(t *testing.T) {

	en, err := ut.GetTranslator("en")
	Equal(t, err, nil)

	result, err := en.FmtCurrencySafe(ut.CurrencyStandard, "USD", 12345.6789)
	Equal(t, err, nil)
	Equal(t, result, "$12,345.68")

	result = en.FmtCurrency(ut.CurrencyStandard, "USD", 12345.6789)
	Equal(t, result, "$12,345.68")

	result, err = en.FmtCurrencySafe(ut.CurrencyAccounting, "USD", 12345.6789)
	Equal(t, err, nil)
	Equal(t, result, "$12,345.68")

	result = en.FmtCurrency(ut.CurrencyAccounting, "USD", 12345.6789)
	Equal(t, result, "$12,345.68")

	result, err = en.FmtCurrencySafe(ut.CurrencyStandard, "USD", -12345.6789)
	Equal(t, err, nil)
	Equal(t, result, "-$12,345.68")

	result = en.FmtCurrency(ut.CurrencyStandard, "USD", -12345.6789)
	Equal(t, result, "-$12,345.68")

	result, err = en.FmtCurrencySafe(ut.CurrencyAccounting, "USD", -12345.6789)
	Equal(t, err, nil)
	Equal(t, result, "($12,345.68)")

	result = en.FmtCurrency(ut.CurrencyAccounting, "USD", -12345.6789)
	Equal(t, result, "($12,345.68)")

	result, err = en.FmtCurrencySafe(ut.CurrencyStandard, "WHAT???", 12345.6789)
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "**** WARNING **** unknown currency: WHAT???")

	result = en.FmtCurrency(ut.CurrencyStandard, "WHAT???", 12345.6789)
	Equal(t, result, "WHAT???12,345.68")

	result, err = en.FmtCurrencySafe(ut.CurrencyAccounting, "WHAT???", 12345.6789)
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "**** WARNING **** unknown currency: WHAT???")

	result = en.FmtCurrency(ut.CurrencyAccounting, "WHAT???", 12345.6789)
	Equal(t, result, "WHAT???12,345.68")

	// try some really big numbers to make sure weird floaty stuff doesn't
	// happen
	result, err = en.FmtCurrencySafe(ut.CurrencyStandard, "USD", 12345000000000.6789)
	Equal(t, err, nil)
	Equal(t, result, "$12,345,000,000,000.68")

	result = en.FmtCurrency(ut.CurrencyStandard, "USD", 12345000000000.6789)
	Equal(t, result, "$12,345,000,000,000.68")

	result, err = en.FmtCurrencySafe(ut.CurrencyAccounting, "USD", 12345000000000.6789)
	Equal(t, err, nil)
	Equal(t, result, "$12,345,000,000,000.68")

	result = en.FmtCurrency(ut.CurrencyAccounting, "USD", 12345000000000.6789)
	Equal(t, result, "$12,345,000,000,000.68")

	result, err = en.FmtCurrencySafe(ut.CurrencyStandard, "USD", -12345000000000.6789)
	Equal(t, err, nil)
	Equal(t, result, "-$12,345,000,000,000.68")

	result = en.FmtCurrency(ut.CurrencyStandard, "USD", -12345000000000.6789)
	Equal(t, result, "-$12,345,000,000,000.68")

	result, err = en.FmtCurrencySafe(ut.CurrencyAccounting, "USD", -12345000000000.6789)
	Equal(t, err, nil)
	Equal(t, result, "($12,345,000,000,000.68)")

	result = en.FmtCurrency(ut.CurrencyAccounting, "USD", -12345000000000.6789)
	Equal(t, result, "($12,345,000,000,000.68)")

	saq, err := ut.GetTranslator("saq")
	Equal(t, err, nil)

	// no Formatting or symbols for "saq" in locales
	result, err = saq.FmtCurrencySafe(ut.CurrencyStandard, "USD", 12345.6789)
	Equal(t, err, nil)
	Equal(t, result, "1234568")

	result = saq.FmtCurrency(ut.CurrencyStandard, "USD", 12345.6789)
	Equal(t, result, "1234568")

	result, err = saq.FmtCurrencySafe(ut.CurrencyAccounting, "USD", 12345.6789)
	Equal(t, err, nil)
	Equal(t, result, "1234568")

	result = saq.FmtCurrency(ut.CurrencyAccounting, "USD", 12345.6789)
	Equal(t, result, "1234568")

	ar, err := ut.GetTranslator("ar")
	Equal(t, err, nil)

	result, err = ar.FmtCurrencySafe(ut.CurrencyStandard, "USD", -12345.6789)
	Equal(t, err, nil)
	Equal(t, result, "\u200F-US$\u00A012٬345٫68")

	result = ar.FmtCurrency(ut.CurrencyStandard, "USD", -12345.6789)
	Equal(t, result, "\u200F-US$\u00A012٬345٫68")

	result, err = ar.FmtCurrencySafe(ut.CurrencyAccounting, "USD", -12345.6789)
	Equal(t, err, nil)
	Equal(t, result, "\u200F-US$\u00A012٬345٫68")

	result = ar.FmtCurrency(ut.CurrencyAccounting, "USD", -12345.6789)
	Equal(t, result, "\u200F-US$\u00A012٬345٫68")

	// And one more for with some unusual symbols for good measure
	result, err = en.FmtCurrencySafe(ut.CurrencyStandard, "USD", 0.0084)
	Equal(t, err, nil)
	Equal(t, result, "$0.01")

	result = en.FmtCurrency(ut.CurrencyStandard, "USD", 0.0084)
	Equal(t, result, "$0.01")
}

func TestCurrencyWholeEn(t *testing.T) {

	en, err := ut.GetTranslator("en")
	Equal(t, err, nil)

	result, err := en.FmtCurrencyWholeSafe(ut.CurrencyStandard, "USD", 12345.6789)
	Equal(t, err, nil)
	Equal(t, result, "$12,346")

	result = en.FmtCurrencyWhole(ut.CurrencyStandard, "USD", 12345.6789)
	Equal(t, result, "$12,346")

	result, err = en.FmtCurrencyWholeSafe(ut.CurrencyAccounting, "USD", 12345.6789)
	Equal(t, err, nil)
	Equal(t, result, "$12,346")

	result = en.FmtCurrencyWhole(ut.CurrencyAccounting, "USD", 12345.6789)
	Equal(t, result, "$12,346")

	result, err = en.FmtCurrencyWholeSafe(ut.CurrencyStandard, "USD", -12345.6789)
	Equal(t, err, nil)
	Equal(t, result, "-$12,346")

	result = en.FmtCurrencyWhole(ut.CurrencyStandard, "USD", -12345.6789)
	Equal(t, result, "-$12,346")

	result, err = en.FmtCurrencyWholeSafe(ut.CurrencyAccounting, "USD", -12345.6789)
	Equal(t, err, nil)
	Equal(t, result, "($12,346)")

	result = en.FmtCurrencyWhole(ut.CurrencyAccounting, "USD", -12345.6789)
	Equal(t, result, "($12,346)")

	result, err = en.FmtCurrencyWholeSafe(ut.CurrencyStandard, "WHAT???", 12345.6789)
	NotEqual(t, err, nil)
	Equal(t, result, "WHAT???12,346")

	result = en.FmtCurrencyWhole(ut.CurrencyStandard, "WHAT???", 12345.6789)
	Equal(t, result, "WHAT???12,346")

	// try some really big numbers to make sure weird floaty stuff doesn't
	// happen
	result, err = en.FmtCurrencyWholeSafe(ut.CurrencyStandard, "USD", 12345000000000.6789)
	Equal(t, err, nil)
	Equal(t, result, "$12,345,000,000,001")

	result, err = en.FmtCurrencyWholeSafe(ut.CurrencyAccounting, "USD", 12345000000000.6789)
	Equal(t, err, nil)
	Equal(t, result, "$12,345,000,000,001")

	result, err = en.FmtCurrencyWholeSafe(ut.CurrencyStandard, "USD", -12345000000000.6789)
	Equal(t, err, nil)
	Equal(t, result, "-$12,345,000,000,001")

	result = en.FmtCurrencyWhole(ut.CurrencyStandard, "USD", -12345000000000.6789)
	Equal(t, result, "-$12,345,000,000,001")

	result, err = en.FmtCurrencyWholeSafe(ut.CurrencyAccounting, "USD", -12345000000000.6789)
	Equal(t, err, nil)
	Equal(t, result, "($12,345,000,000,001)")

	result = en.FmtCurrencyWhole(ut.CurrencyAccounting, "USD", -12345000000000.6789)
	Equal(t, result, "($12,345,000,000,001)")
}

func TestNumberEn(t *testing.T) {

	en, err := ut.GetTranslator("en")
	Equal(t, err, nil)

	// check basic english
	result := en.FmtNumber(12345.6789)
	Equal(t, result, "12,345.679")

	result = en.FmtNumber(-12345.6789)
	Equal(t, result, "-12,345.679")

	result = en.FmtNumber(123456789)
	Equal(t, result, "123,456,789")

	hi, err := ut.GetTranslator("hi")
	Equal(t, err, nil)

	// check Hindi - different group sizes
	result = hi.FmtNumber(12345.6789)
	Equal(t, result, "12,345.679")

	result = hi.FmtNumber(-12345.6789)
	Equal(t, result, "-12,345.679")

	result = hi.FmtNumber(123456789)
	Equal(t, result, "12,34,56,789")

	uz, err := ut.GetTranslator("uz")
	Equal(t, err, nil)

	// check Uzbek - something with a partial fallback
	result = uz.FmtNumber(12345.6789)
	Equal(t, result, "12٬345٫679")

	result = uz.FmtNumber(-12345.6789)
	Equal(t, result, "-12٬345٫679")

	result = uz.FmtNumber(123456789)
	Equal(t, result, "123٬456٬789")
}

func TestNumberWholeEn(t *testing.T) {

	en, err := ut.GetTranslator("en")
	Equal(t, err, nil)

	result := en.FmtNumberWhole(12345.6789)
	Equal(t, result, "12,346")

	result = en.FmtNumberWhole(-12345.6789)
	Equal(t, result, "-12,346")
}

func TestPercentEn(t *testing.T) {

	en, err := ut.GetTranslator("en")
	Equal(t, err, nil)

	result := en.FmtPercent(0.01234)
	Equal(t, result, "1%")

	result = en.FmtPercent(0.1234)
	Equal(t, result, "12%")

	result = en.FmtPercent(1.234)
	Equal(t, result, "123%")

	result = en.FmtPercent(12.34)
	Equal(t, result, "1,234%")
}
