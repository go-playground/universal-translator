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

func TestDateTimeTh(t *testing.T) {
	th, err := ut.GetTranslator("th")
	Equal(t, err, nil)

	datetime, err := time.Parse(dateTimeString, dateTimeString)
	Equal(t, err, nil)

	// test the public method
	dt, err := th.FmtDateFullSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "วันจันทร์ที่ 2 มกราคม คริสต์ศักราช 2006")

	dt, err = th.FmtDateFullSafe(bc300DateTime)
	Equal(t, err, nil)
	Equal(t, dt, "วันเสาร์ที่ 2 มกราคม 300 ปีก่อนคริสต์ศักราช")

	dt = th.FmtDateFull(datetime)
	Equal(t, dt, "วันจันทร์ที่ 2 มกราคม คริสต์ศักราช 2006")

	dt = th.FmtDateFull(bc300DateTime)
	Equal(t, dt, "วันเสาร์ที่ 2 มกราคม 300 ปีก่อนคริสต์ศักราช")

	dt, err = th.FmtDateLongSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "2 มกราคม ค.ศ. 2006")

	dt, err = th.FmtDateLongSafe(bc300DateTime)
	Equal(t, err, nil)
	Equal(t, dt, "2 มกราคม 300 ปีก่อน ค.ศ.")

	dt = th.FmtDateLong(datetime)
	Equal(t, dt, "2 มกราคม ค.ศ. 2006")

	dt = th.FmtDateLong(bc300DateTime)
	Equal(t, dt, "2 มกราคม 300 ปีก่อน ค.ศ.")

	dt, err = th.FmtDateMediumSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "2 ม.ค. 2006")

	dt = th.FmtDateMedium(datetime)
	Equal(t, dt, "2 ม.ค. 2006")

	dt, err = th.FmtDateShortSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "2/1/06")

	dt = th.FmtDateShort(datetime)
	Equal(t, dt, "2/1/06")

	dt, err = th.FmtTimeFullSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "15 นาฬิกา 04 นาที 05 วินาที")

	dt = th.FmtTimeFull(datetime)
	Equal(t, dt, "15 นาฬิกา 04 นาที 05 วินาที")

	dt, err = th.FmtTimeLongSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "15 นาฬิกา 04 นาที 05 วินาที")

	dt = th.FmtTimeLong(datetime)
	Equal(t, dt, "15 นาฬิกา 04 นาที 05 วินาที")

	dt, err = th.FmtTimeMediumSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "15:04:05")

	dt = th.FmtTimeMedium(datetime)
	Equal(t, dt, "15:04:05")

	dt, err = th.FmtTimeShortSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "15:04")

	dt = th.FmtTimeShort(datetime)
	Equal(t, dt, "15:04")

	dt, err = th.FmtDateTimeFullSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "วันจันทร์ที่ 2 มกราคม คริสต์ศักราช 2006 15 นาฬิกา 04 นาที 05 วินาที")

	dt, err = th.FmtDateTimeFullSafe(bc300DateTime)
	Equal(t, err, nil)
	Equal(t, dt, "วันเสาร์ที่ 2 มกราคม 300 ปีก่อนคริสต์ศักราช 3 นาฬิกา 04 นาที 05 วินาที")

	dt = th.FmtDateTimeFull(datetime)
	Equal(t, dt, "วันจันทร์ที่ 2 มกราคม คริสต์ศักราช 2006 15 นาฬิกา 04 นาที 05 วินาที")

	dt = th.FmtDateTimeFull(bc300DateTime)
	Equal(t, dt, "วันเสาร์ที่ 2 มกราคม 300 ปีก่อนคริสต์ศักราช 3 นาฬิกา 04 นาที 05 วินาที")

	dt, err = th.FmtDateTimeLongSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "2 มกราคม ค.ศ. 2006 15 นาฬิกา 04 นาที 05 วินาที")

	dt, err = th.FmtDateTimeLongSafe(bc300DateTime)
	Equal(t, err, nil)
	Equal(t, dt, "2 มกราคม 300 ปีก่อน ค.ศ. 3 นาฬิกา 04 นาที 05 วินาที")

	dt = th.FmtDateTimeLong(datetime)
	Equal(t, dt, "2 มกราคม ค.ศ. 2006 15 นาฬิกา 04 นาที 05 วินาที")

	dt = th.FmtDateTimeLong(bc300DateTime)
	Equal(t, dt, "2 มกราคม 300 ปีก่อน ค.ศ. 3 นาฬิกา 04 นาที 05 วินาที")

	dt, err = th.FmtDateTimeMediumSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "2 ม.ค. 2006 15:04:05")

	dt = th.FmtDateTimeMedium(datetime)
	Equal(t, dt, "2 ม.ค. 2006 15:04:05")

	dt, err = th.FmtDateTimeShortSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "2/1/06 15:04")

	dt = th.FmtDateTimeShort(datetime)
	Equal(t, dt, "2/1/06 15:04")

	dt, err = th.FmtDateTimeSafe(datetime, "MMMM d yy")
	Equal(t, err, nil)
	Equal(t, dt, "มกราคม 2 06")

	dt = th.FmtDateTime(datetime, "MMMM d yy")
	Equal(t, err, nil)
	Equal(t, dt, "มกราคม 2 06")

	dt, err = th.FmtDateTimeSafe(datetime, "not a date pattern")
	NotEqual(t, err, nil)
	Equal(t, dt, "")
	Equal(t, err.Error(), "unknown datetime format unit: n")
}

func TestCurrencyTh(t *testing.T) {

	th, err := ut.GetTranslator("th")
	Equal(t, err, nil)

	result, err := th.FmtCurrencySafe(ut.CurrencyStandard, "THB", 12345.6789)
	Equal(t, err, nil)
	Equal(t, result, "THB12,345.68")

	result = th.FmtCurrency(ut.CurrencyStandard, "THB", 12345.6789)
	Equal(t, result, "THB12,345.68")

	result, err = th.FmtCurrencySafe(ut.CurrencyAccounting, "THB", 12345.6789)
	Equal(t, err, nil)
	Equal(t, result, "THB12,345.68")

	result = th.FmtCurrency(ut.CurrencyAccounting, "THB", 12345.6789)
	Equal(t, result, "THB12,345.68")

	result, err = th.FmtCurrencySafe(ut.CurrencyStandard, "THB", -12345.6789)
	Equal(t, err, nil)
	Equal(t, result, "-THB12,345.68")

	result = th.FmtCurrency(ut.CurrencyStandard, "THB", -12345.6789)
	Equal(t, result, "-THB12,345.68")

	result, err = th.FmtCurrencySafe(ut.CurrencyAccounting, "THB", -12345.6789)
	Equal(t, err, nil)
	Equal(t, result, "(THB12,345.68)")

	result = th.FmtCurrency(ut.CurrencyAccounting, "THB", -12345.6789)
	Equal(t, result, "(THB12,345.68)")

	result, err = th.FmtCurrencySafe(ut.CurrencyStandard, "WHAT???", 12345.6789)
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "**** WARNING **** unknown currency: WHAT???")

	result = th.FmtCurrency(ut.CurrencyStandard, "WHAT???", 12345.6789)
	Equal(t, result, "WHAT???12,345.68")

	result, err = th.FmtCurrencySafe(ut.CurrencyAccounting, "WHAT???", 12345.6789)
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "**** WARNING **** unknown currency: WHAT???")

	result = th.FmtCurrency(ut.CurrencyAccounting, "WHAT???", 12345.6789)
	Equal(t, result, "WHAT???12,345.68")

	// try some really big numbers to make sure weird floaty stuff doesn't
	// happen
	result, err = th.FmtCurrencySafe(ut.CurrencyStandard, "THB", 12345000000000.6789)
	Equal(t, err, nil)
	Equal(t, result, "THB12,345,000,000,000.68")

	result = th.FmtCurrency(ut.CurrencyStandard, "THB", 12345000000000.6789)
	Equal(t, result, "THB12,345,000,000,000.68")

	result, err = th.FmtCurrencySafe(ut.CurrencyAccounting, "THB", 12345000000000.6789)
	Equal(t, err, nil)
	Equal(t, result, "THB12,345,000,000,000.68")

	result = th.FmtCurrency(ut.CurrencyAccounting, "THB", 12345000000000.6789)
	Equal(t, result, "THB12,345,000,000,000.68")

	result, err = th.FmtCurrencySafe(ut.CurrencyStandard, "THB", -12345000000000.6789)
	Equal(t, err, nil)
	Equal(t, result, "-THB12,345,000,000,000.68")

	result = th.FmtCurrency(ut.CurrencyStandard, "THB", -12345000000000.6789)
	Equal(t, result, "-THB12,345,000,000,000.68")

	result, err = th.FmtCurrencySafe(ut.CurrencyAccounting, "THB", -12345000000000.6789)
	Equal(t, err, nil)
	Equal(t, result, "(THB12,345,000,000,000.68)")

	result = th.FmtCurrency(ut.CurrencyAccounting, "THB", -12345000000000.6789)
	Equal(t, result, "(THB12,345,000,000,000.68)")

	en, err := ut.GetTranslator("en")
	Equal(t, err, nil)

	result, err = en.FmtCurrencySafe(ut.CurrencyStandard, "THB", 12345.6789)
	Equal(t, err, nil)
	Equal(t, result, "12,345.68")

	result = en.FmtCurrency(ut.CurrencyStandard, "THB", 12345.6789)
	Equal(t, result, "12,345.68")

	result, err = en.FmtCurrencySafe(ut.CurrencyAccounting, "THB", 12345.6789)
	Equal(t, err, nil)
	Equal(t, result, "12,345.68")

	result = en.FmtCurrency(ut.CurrencyAccounting, "THB", 12345.6789)
	Equal(t, result, "12,345.68")

	ar, err := ut.GetTranslator("ar")
	Equal(t, err, nil)

	result, err = ar.FmtCurrencySafe(ut.CurrencyStandard, "THB", -12345.6789)
	Equal(t, err, nil)
	Equal(t, result, "\u200F-฿\u00A012٬345٫68")

	result = ar.FmtCurrency(ut.CurrencyStandard, "THB", -12345.6789)
	Equal(t, result, "\u200F-฿\u00A012٬345٫68")

	result, err = ar.FmtCurrencySafe(ut.CurrencyAccounting, "THB", -12345.6789)
	Equal(t, err, nil)
	Equal(t, result, "\u200F-฿\u00A012٬345٫68")

	result = ar.FmtCurrency(ut.CurrencyAccounting, "THB", -12345.6789)
	Equal(t, result, "\u200F-฿\u00A012٬345٫68")

	// And one more for with some unusual symbols for good measure
	result, err = th.FmtCurrencySafe(ut.CurrencyStandard, "THB", 0.0084)
	Equal(t, err, nil)
	Equal(t, result, "THB0.01")

	result = th.FmtCurrency(ut.CurrencyStandard, "THB", 0.0084)
	Equal(t, result, "THB0.01")
}

func TestCurrencyWholeTh(t *testing.T) {

	th, err := ut.GetTranslator("th")
	Equal(t, err, nil)

	result, err := th.FmtCurrencyWholeSafe(ut.CurrencyStandard, "THB", 12345.6789)
	Equal(t, err, nil)
	Equal(t, result, "THB12,346")

	result = th.FmtCurrencyWhole(ut.CurrencyStandard, "THB", 12345.6789)
	Equal(t, result, "THB12,346")

	result, err = th.FmtCurrencyWholeSafe(ut.CurrencyAccounting, "THB", 12345.6789)
	Equal(t, err, nil)
	Equal(t, result, "THB12,346")

	result = th.FmtCurrencyWhole(ut.CurrencyAccounting, "THB", 12345.6789)
	Equal(t, result, "THB12,346")

	result, err = th.FmtCurrencyWholeSafe(ut.CurrencyStandard, "THB", -12345.6789)
	Equal(t, err, nil)
	Equal(t, result, "-THB12,346")

	result = th.FmtCurrencyWhole(ut.CurrencyStandard, "THB", -12345.6789)
	Equal(t, result, "-THB12,346")

	result, err = th.FmtCurrencyWholeSafe(ut.CurrencyAccounting, "THB", -12345.6789)
	Equal(t, err, nil)
	Equal(t, result, "(THB12,346)")

	result = th.FmtCurrencyWhole(ut.CurrencyAccounting, "THB", -12345.6789)
	Equal(t, result, "(THB12,346)")

	result, err = th.FmtCurrencyWholeSafe(ut.CurrencyStandard, "WHAT???", 12345.6789)
	NotEqual(t, err, nil)
	Equal(t, result, "WHAT???12,346")

	result = th.FmtCurrencyWhole(ut.CurrencyStandard, "WHAT???", 12345.6789)
	Equal(t, result, "WHAT???12,346")

	// try some really big numbers to make sure weird floaty stuff doesn't
	// happen
	result, err = th.FmtCurrencyWholeSafe(ut.CurrencyStandard, "THB", 12345000000000.6789)
	Equal(t, err, nil)
	Equal(t, result, "THB12,345,000,000,001")

	result, err = th.FmtCurrencyWholeSafe(ut.CurrencyAccounting, "THB", 12345000000000.6789)
	Equal(t, err, nil)
	Equal(t, result, "THB12,345,000,000,001")

	result, err = th.FmtCurrencyWholeSafe(ut.CurrencyStandard, "THB", -12345000000000.6789)
	Equal(t, err, nil)
	Equal(t, result, "-THB12,345,000,000,001")

	result = th.FmtCurrencyWhole(ut.CurrencyStandard, "THB", -12345000000000.6789)
	Equal(t, result, "-THB12,345,000,000,001")

	result, err = th.FmtCurrencyWholeSafe(ut.CurrencyAccounting, "THB", -12345000000000.6789)
	Equal(t, err, nil)
	Equal(t, result, "(THB12,345,000,000,001)")

	result = th.FmtCurrencyWhole(ut.CurrencyAccounting, "THB", -12345000000000.6789)
	Equal(t, result, "(THB12,345,000,000,001)")
}

func TestNumberTh(t *testing.T) {

	th, err := ut.GetTranslator("th")
	Equal(t, err, nil)

	// check basic english
	result := th.FmtNumber(12345.6789)
	Equal(t, result, "12,345.679")

	result = th.FmtNumber(-12345.6789)
	Equal(t, result, "-12,345.679")

	result = th.FmtNumber(123456789)
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

func TestNumberWholeTh(t *testing.T) {

	th, err := ut.GetTranslator("th")
	Equal(t, err, nil)

	result := th.FmtNumberWhole(12345.6789)
	Equal(t, result, "12,346")

	result = th.FmtNumberWhole(-12345.6789)
	Equal(t, result, "-12,346")
}

func TestPercentTh(t *testing.T) {

	th, err := ut.GetTranslator("th")
	Equal(t, err, nil)

	result := th.FmtPercent(0.01234)
	Equal(t, result, "1%")

	result = th.FmtPercent(0.1234)
	Equal(t, result, "12%")

	result = th.FmtPercent(1.234)
	Equal(t, result, "123%")

	result = th.FmtPercent(12.34)
	Equal(t, result, "1,234%")
}
