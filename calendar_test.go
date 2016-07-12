package ut_test

import "time"

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
