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

var (
	dateTimeString = "Jan 2, 2006 at 3:04:05pm"
	bc300DateTime  = time.Date(-300, 1, 2, 3, 4, 5, 0, time.UTC)
)

func TestDateTimeEn(t *testing.T) {

	en, err := ut.GetTranslator("en")
	Equal(t, err, nil)

	datetime, err := time.Parse(dateTimeString, dateTimeString)
	Equal(t, err, nil)

	// test the public method
	dt, err := en.FmtDateFullSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "Monday, January 2, 2006")

	dt = en.FmtDateFull(datetime)
	Equal(t, dt, "Monday, January 2, 2006")

	dt, err = en.FmtDateLongSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "January 2, 2006")

	dt = en.FmtDateLong(datetime)
	Equal(t, dt, "January 2, 2006")

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

	dt = en.FmtDateTimeFull(datetime)
	Equal(t, dt, "Monday, January 2, 2006 at 3:04:05 PM")

	dt, err = en.FmtDateTimeLongSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "January 2, 2006 at 3:04:05 PM")

	dt = en.FmtDateTimeLong(datetime)
	Equal(t, dt, "January 2, 2006 at 3:04:05 PM")

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

func TestDateTimeTh(t *testing.T) {
	th, err := ut.GetTranslator("th")
	Equal(t, err, nil)

	datetime, err := time.Parse(dateTimeString, dateTimeString)
	Equal(t, err, nil)

	// test the public method
	dt, err := th.FmtDateFullSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "วันจันทร์ที่ 2 มกราคม 2006 คริสต์ศักราช")

	dt, err = th.FmtDateFullSafe(bc300DateTime)
	Equal(t, err, nil)
	Equal(t, dt, "วันเสาร์ที่ 2 มกราคม ปีก่อนคริสต์ศักราช 300")

	dt = th.FmtDateFull(datetime)
	Equal(t, dt, "วันจันทร์ที่ 2 มกราคม 2006 คริสต์ศักราช")

	dt = th.FmtDateFull(bc300DateTime)
	Equal(t, dt, "วันเสาร์ที่ 2 มกราคม ปีก่อนคริสต์ศักราช 300")

	dt, err = th.FmtDateLongSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "2 มกราคม 2006 ค.ศ.")

	dt, err = th.FmtDateLongSafe(bc300DateTime)
	Equal(t, err, nil)
	Equal(t, dt, "2 มกราคม ปีก่อน ค.ศ. 300")

	dt = th.FmtDateLong(datetime)
	Equal(t, dt, "2 มกราคม 2006 ค.ศ.")

	dt = th.FmtDateLong(bc300DateTime)
	Equal(t, dt, "2 มกราคม ปีก่อน ค.ศ. 300")

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
	Equal(t, dt, "วันจันทร์ที่ 2 มกราคม 2006 คริสต์ศักราช 15 นาฬิกา 04 นาที 05 วินาที")

	dt, err = th.FmtDateTimeFullSafe(bc300DateTime)
	Equal(t, err, nil)
	Equal(t, dt, "วันเสาร์ที่ 2 มกราคม ปีก่อนคริสต์ศักราช 300 3 นาฬิกา 04 นาที 05 วินาที")

	dt = th.FmtDateTimeFull(datetime)
	Equal(t, dt, "วันจันทร์ที่ 2 มกราคม 2006 คริสต์ศักราช 15 นาฬิกา 04 นาที 05 วินาที")

	dt = th.FmtDateTimeFull(bc300DateTime)
	Equal(t, dt, "วันเสาร์ที่ 2 มกราคม ปีก่อนคริสต์ศักราช 300 3 นาฬิกา 04 นาที 05 วินาที")

	dt, err = th.FmtDateTimeLongSafe(datetime)
	Equal(t, err, nil)
	Equal(t, dt, "2 มกราคม 2006 ค.ศ. 15 นาฬิกา 04 นาที 05 วินาที")

	dt, err = th.FmtDateTimeLongSafe(bc300DateTime)
	Equal(t, err, nil)
	Equal(t, dt, "2 มกราคม ปีก่อน ค.ศ. 300 3 นาฬิกา 04 นาที 05 วินาที")

	dt = th.FmtDateTimeLong(datetime)
	Equal(t, dt, "2 มกราคม 2006 ค.ศ. 15 นาฬิกา 04 นาที 05 วินาที")

	dt = th.FmtDateTimeLong(bc300DateTime)
	Equal(t, dt, "2 มกราคม ปีก่อน ค.ศ. 300 3 นาฬิกา 04 นาที 05 วินาที")

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
