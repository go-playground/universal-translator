package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"strings"
	"sync"
	"text/template"

	"golang.org/x/text/unicode/cldr"

	i18n "github.com/go-playground/universal-translator"
	"github.com/kr/pretty"
)

// numbers:
//   symbols:
//     decimal: .
//     group: ','
//     negative: '-'
//     percent: '%'
//     permille: "\u2030"
//   formats:
//     decimal: '#,##0.###'
//     currency: "\xA4#,##0.00;(\xA4#,##0.00)"
//     percent: '#,##0%'
//   currencies:
//     JPY:
//       symbol: "\xA5"
//     USD:
//       symbol: $

func main() {
	var decoder cldr.Decoder
	cldr, err := decoder.DecodePath("data/core")
	if err != nil {
		panic(err)
	}

	locs := map[string]string{}
	numbers := map[string]i18n.Number{}
	calendars := map[string]i18n.Calendar{}
	locales := cldr.Locales()
	for _, loc := range locales {

		ldml := cldr.RawLDML(loc)
		if ldml.Numbers == nil {
			continue
		}
		var number i18n.Number
		if len(ldml.Numbers.Symbols) > 0 {
			symbol := ldml.Numbers.Symbols[0]
			if len(symbol.Decimal) > 0 {
				number.Symbols.Decimal = symbol.Decimal[0].Data()
			}
			if len(symbol.Group) > 0 {
				number.Symbols.Group = symbol.Group[0].Data()
			}
			if len(symbol.MinusSign) > 0 {
				number.Symbols.Negative = symbol.MinusSign[0].Data()
			}
			if len(symbol.PercentSign) > 0 {
				number.Symbols.Percent = symbol.PercentSign[0].Data()
			}
			if len(symbol.PerMille) > 0 {
				number.Symbols.PerMille = symbol.PerMille[0].Data()
			}
		}
		if len(ldml.Numbers.DecimalFormats) > 0 && len(ldml.Numbers.DecimalFormats[0].DecimalFormatLength) > 0 {
			number.Formats.Decimal = ldml.Numbers.DecimalFormats[0].DecimalFormatLength[0].DecimalFormat[0].Pattern[0].Data()
		}
		if len(ldml.Numbers.CurrencyFormats) > 0 && len(ldml.Numbers.CurrencyFormats[0].CurrencyFormatLength) > 0 {
			number.Formats.Currency = ldml.Numbers.CurrencyFormats[0].CurrencyFormatLength[0].CurrencyFormat[0].Pattern[0].Data()
		}
		if len(ldml.Numbers.PercentFormats) > 0 && len(ldml.Numbers.PercentFormats[0].PercentFormatLength) > 0 {
			number.Formats.Percent = ldml.Numbers.PercentFormats[0].PercentFormatLength[0].PercentFormat[0].Pattern[0].Data()
		}
		if ldml.Numbers.Currencies != nil {
			for _, currency := range ldml.Numbers.Currencies.Currency {
				var c i18n.Currency
				c.Currency = currency.Type
				if len(currency.DisplayName) > 0 {
					c.DisplayName = currency.DisplayName[0].Data()
				}
				if len(currency.Symbol) > 0 {
					c.Symbol = currency.Symbol[0].Data()
				}
				number.Currencies = append(number.Currencies, c)
			}
		}
		numbers[loc] = number
		locs[loc] = strings.Replace(loc, "_", "", -1)

		if ldml.Dates != nil && ldml.Dates.Calendars != nil {
			var calendar i18n.Calendar
			ldmlCar := ldml.Dates.Calendars.Calendar[0]
			for _, cal := range ldml.Dates.Calendars.Calendar {
				if cal.Type == "gregorian" {
					ldmlCar = cal
				}
			}
			if ldmlCar.DateFormats != nil {
				for _, datefmt := range ldmlCar.DateFormats.DateFormatLength {
					switch datefmt.Type {
					case "full":
						calendar.Formats.Date.Full = datefmt.DateFormat[0].Pattern[0].Data()
					case "long":
						calendar.Formats.Date.Long = datefmt.DateFormat[0].Pattern[0].Data()
					case "medium":
						calendar.Formats.Date.Medium = datefmt.DateFormat[0].Pattern[0].Data()
					case "short":
						calendar.Formats.Date.Short = datefmt.DateFormat[0].Pattern[0].Data()
					}
				}
			}

			if ldmlCar.TimeFormats != nil {
				for _, datefmt := range ldmlCar.TimeFormats.TimeFormatLength {
					switch datefmt.Type {
					case "full":
						calendar.Formats.Time.Full = datefmt.TimeFormat[0].Pattern[0].Data()
					case "long":
						calendar.Formats.Time.Long = datefmt.TimeFormat[0].Pattern[0].Data()
					case "medium":
						calendar.Formats.Time.Medium = datefmt.TimeFormat[0].Pattern[0].Data()
					case "short":
						calendar.Formats.Time.Short = datefmt.TimeFormat[0].Pattern[0].Data()
					}
				}
			}
			if ldmlCar.DateTimeFormats != nil {
				for _, datefmt := range ldmlCar.DateTimeFormats.DateTimeFormatLength {
					switch datefmt.Type {
					case "full":
						calendar.Formats.DateTime.Full = datefmt.DateTimeFormat[0].Pattern[0].Data()
					case "long":
						calendar.Formats.DateTime.Long = datefmt.DateTimeFormat[0].Pattern[0].Data()
					case "medium":
						calendar.Formats.DateTime.Medium = datefmt.DateTimeFormat[0].Pattern[0].Data()
					case "short":
						calendar.Formats.DateTime.Short = datefmt.DateTimeFormat[0].Pattern[0].Data()
					}
				}
			}
			if ldmlCar.Months != nil {
				for _, monthctx := range ldmlCar.Months.MonthContext {
					for _, months := range monthctx.MonthWidth {
						var i18nMonth i18n.CalendarMonthFormatNameValue
						for _, m := range months.Month {
							switch m.Type {
							case "1":
								i18nMonth.Jan = m.Data()
							case "2":
								i18nMonth.Feb = m.Data()
							case "3":
								i18nMonth.Mar = m.Data()
							case "4":
								i18nMonth.Apr = m.Data()
							case "5":
								i18nMonth.May = m.Data()
							case "6":
								i18nMonth.Jun = m.Data()
							case "7":
								i18nMonth.Jul = m.Data()
							case "8":
								i18nMonth.Aug = m.Data()
							case "9":
								i18nMonth.Sep = m.Data()
							case "10":
								i18nMonth.Oct = m.Data()
							case "11":
								i18nMonth.Nov = m.Data()
							case "12":
								i18nMonth.Dec = m.Data()
							}
						}
						switch months.Type {
						case "abbreviated":
							calendar.FormatNames.Months.Abbreviated = i18nMonth
						case "narrow":
							calendar.FormatNames.Months.Narrow = i18nMonth
						case "short":
							calendar.FormatNames.Months.Short = i18nMonth
						case "wide":
							calendar.FormatNames.Months.Wide = i18nMonth
						}
					}
				}
			}
			if ldmlCar.Days != nil {
				for _, dayctx := range ldmlCar.Days.DayContext {
					for _, days := range dayctx.DayWidth {
						var i18nDay i18n.CalendarDayFormatNameValue
						for _, d := range days.Day {
							switch d.Type {
							case "sun":
								i18nDay.Sun = d.Data()
							case "mon":
								i18nDay.Mon = d.Data()
							case "tue":
								i18nDay.Tue = d.Data()
							case "wed":
								i18nDay.Wed = d.Data()
							case "thu":
								i18nDay.Thu = d.Data()
							case "fri":
								i18nDay.Fri = d.Data()
							case "sat":
								i18nDay.Sat = d.Data()
							}
						}
						switch days.Type {
						case "abbreviated":
							calendar.FormatNames.Days.Abbreviated = i18nDay
						case "narrow":
							calendar.FormatNames.Days.Narrow = i18nDay
						case "short":
							calendar.FormatNames.Days.Short = i18nDay
						case "wide":
							calendar.FormatNames.Days.Wide = i18nDay
						}
					}
				}
			}
			if ldmlCar.DayPeriods != nil {
				for _, ctx := range ldmlCar.DayPeriods.DayPeriodContext {
					for _, width := range ctx.DayPeriodWidth {
						var i18nPeriod i18n.CalendarPeriodFormatNameValue
						for _, d := range width.DayPeriod {
							switch d.Type {
							case "am":
								if i18nPeriod.AM == "" {
									i18nPeriod.AM = d.Data()
								}
							case "pm":
								if i18nPeriod.PM == "" {
									i18nPeriod.PM = d.Data()
								}
							}
						}
						switch width.Type {
						case "abbreviated":
							calendar.FormatNames.Periods.Abbreviated = i18nPeriod
						case "narrow":
							calendar.FormatNames.Periods.Narrow = i18nPeriod
						case "short":
							calendar.FormatNames.Periods.Short = i18nPeriod
						case "wide":
							calendar.FormatNames.Periods.Wide = i18nPeriod
						}
					}
				}
				// var empty i18n.CalendarPeriodFormatNameValue
				// if calendar.FormatNames.Periods.Abbreviated == empty {
				// 	calendar.FormatNames.Periods.Abbreviated = calendar.FormatNames.Periods.Wide
				// }
			}
			calendars[loc] = calendar
		}
	}

	var wg sync.WaitGroup
	wg.Add(len(numbers))
	for locale, number := range numbers {
		go func(locale string, number i18n.Number) {

			localeNoUnderscore := strings.Replace(locale, "_", "", -1)
			defer func() { wg.Done() }()
			path := "../../zzz_generated_" + locale + "_"
			// if _, err := os.Stat(path); err != nil {
			// 	if err = os.MkdirAll(path, 0777); err != nil {
			// 		panic(err)
			// 	}
			// }
			numberFile, err := os.Create(path + "number.go")
			if err != nil {
				panic(err)
			}
			defer func() { numberFile.Close() }()
			mainFile, err := os.Create(path + "main.go")
			if err != nil {
				panic(err)
			}
			defer func() { mainFile.Close() }()
			currencyFile, err := os.Create(path + "currency.go")
			if err != nil {
				panic(err)
			}
			defer func() { currencyFile.Close() }()

			mainCodes, err := format.Source([]byte(fmt.Sprintf(`package ut

			// new returns a new instance of the locale
			func new%s() *Locale {
				return &Locale {
					Locale: %q,
					Number: Number {
						Symbols: new%sSymbols(),
						Formats: new%sFormats(),
						Currencies: new%sCurrencies(),
					},
					Calendar: new%sCalendar(),
					PluralRule: %spluralRule,
				}
			}
		`, localeNoUnderscore, locale, localeNoUnderscore, localeNoUnderscore, localeNoUnderscore, localeNoUnderscore, localeNoUnderscore)))
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(mainFile, "%s", mainCodes)

			numberCodes, err := format.Source([]byte(fmt.Sprintf(`package ut

			func new%sSymbols() Symbols {
				return %#v
			}

			func new%sFormats() NumberFormats {
				return %#v
			}
		`, localeNoUnderscore, pretty.Formatter(number.Symbols), localeNoUnderscore, pretty.Formatter(number.Formats))))
			if err != nil {
				panic(err)
			}

			numberCodes = []byte(strings.Replace(string(numberCodes), "ut.", "", -1))

			fmt.Fprintf(numberFile, "%s", numberCodes)

			currencyCodes, err := format.Source([]byte(fmt.Sprintf(`package ut

			func new%sCurrencies() []Currency {
				return %#v
			}
		`, localeNoUnderscore, pretty.Formatter(number.Currencies))))
			if err != nil {
				panic(err)
			}

			currencyCodes = []byte(strings.Replace(string(currencyCodes), "ut.", "", -1))
			fmt.Fprintf(currencyFile, "%s", currencyCodes)

			calendar := calendars[locale]
			calendarFile, err := os.Create(path + "calendar.go")
			if err != nil {
				panic(err)
			}
			defer func() { calendarFile.Close() }()

			calendarCodes, err := format.Source([]byte(fmt.Sprintf(`package ut

			func new%sCalendar() Calendar {
				return %#v
			}
		`, localeNoUnderscore, pretty.Formatter(calendar))))
			if err != nil {
				panic(err)
			}

			calendarCodes = []byte(strings.Replace(string(calendarCodes), "ut.", "", -1))
			fmt.Fprintf(calendarFile, "%s", calendarCodes)

			pluralFile, err := os.Create(path + "plural.go")
			if err != nil {
				panic(err)
			}
			defer func() { pluralFile.Close() }()

			pluralCodes, err := format.Source([]byte(fmt.Sprintf(`package ut

			var %spluralRule = "1"
		`, localeNoUnderscore)))
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(pluralFile, "%s", pluralCodes)
		}(locale, number)
	}

	wg.Wait()

	// TODO: make switch with all of the locales + function to return new!
	localesFile, err := os.Create("../../zzz_generated_locales.go")
	if err != nil {
		panic(err)
	}
	defer localesFile.Close()

	tmpl, err := template.New("").Parse(`package ut

		import "errors"

		// GetLocale returns the Locale instance associated with
		// the provided locale string, or returns an error when not found.
		func GetLocale(locale string) (*Locale, error) {
			switch locale {
			{{range $locale, $val := .}}case "{{$locale}}":
				return new{{$val}}(), nil
			{{end}}
			default:
				return nil, errors.New("Unknown locale '" + locale + "'")
			}
		}
	`)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, locs); err != nil {
		panic(err)
	}
	allCodes, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	_, err = localesFile.Write(allCodes)
	if err != nil {
		panic(err)
	}
}
