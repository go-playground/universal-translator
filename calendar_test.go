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

var dateTimeString = "Jan 2, 2006 at 3:04:05pm"

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
