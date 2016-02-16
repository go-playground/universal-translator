package ut_test

import (
	"testing"

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

func TestCurrency(t *testing.T) {

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
	Equal(t, result, "")

	result, err = en.FmtCurrencySafe(ut.CurrencyAccounting, "WHAT???", 12345.6789)
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "**** WARNING **** unknown currency: WHAT???")

	result = en.FmtCurrency(ut.CurrencyAccounting, "WHAT???", 12345.6789)
	Equal(t, result, "")

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

	// Try something that needs a partial fallback
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

	// // And one more for with some unusual symbols for good measure
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
