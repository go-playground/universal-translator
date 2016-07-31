package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/go-playground/universal-translator/resources/locales"

	"golang.org/x/text/unicode/cldr"

	"text/template"
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

// type pluralInfo struct {
// 	path   string
// 	locale string
// 	plural string
// }

// type translator struct {
// 	locale string

// }
//

const (
	locDir      = "../../resources/locales/%s"
	locFilename = locDir + "/%s.go"
)

var (
	prVarFuncs = map[string]string{
		"n": `n, err := locales.N(num)
		if err != nil {
			return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue:num, InnerError: err}
		}

		`,
		"i": `i, err := locales.I(num)
		if err != nil {
			return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue:num, InnerError: err}
		}

		`,
		"v": "v := locales.V(num)\n\n",
		"w": `w, err := locales.W(num)
		if err != nil {
			return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue:num, InnerError: err}
		}

		`,
		"f": `f, err := locales.F(num)
		if err != nil {
			return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue:num, InnerError: err}
		}

		`,
		"t": `t, err := locales.T(num)
		if err != nil {
			return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue:num, InnerError: err}
		}

		`,
	}
)

type translator struct {
	Locale       string
	Plurals      string
	CardinalFunc string
}

var tmpl *template.Template

func main() {

	var err error

	// load template
	tmpl, err = template.ParseGlob("*.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	// load CLDR recourses
	var decoder cldr.Decoder
	cldr, err := decoder.DecodePath("data/core")
	if err != nil {
		panic(err)
	}

	// cardinalPlurals := map[string]

	// for _, p := range ldr.Supplemental().Plurals {

	// 	for _, pr := range p.PluralRules {

	// 		fmt.Println(pr.Locales)

	// 		for _, rule := range pr.PluralRule {
	// 			fmt.Println(rule.Count, rule.Common.Data())
	// 		}
	// 	}
	// }

	for _, l := range cldr.Locales() {

		// // work on uk for the moment
		// if l != "uk" && l != "fil" && l != "gd" {
		// if l != "gd" {
		// 	continue
		// }
		fmt.Println(l)

		baseLocale := strings.SplitN(l, "_", 2)[0]

		trans := &translator{
			Locale: l,
		}

		trans.CardinalFunc, trans.Plurals = parseCardinalPluralRuleFunc(cldr, baseLocale)

		// fmt.Println(trans.CardinalFunc, trans.Plurals)
		// cardinalRules := getLocaleCardinalPluralRules(cldr, baseLocale)
		// fmt.Println("CardinalRules:", l, cardinalRules)
		// Start Plural Rules

		// for _, p := range cldr.Supplemental().Plurals {

		// 	for _, pr := range p.PluralRules {

		// 		locales := strings.Split(pr.Locales, " ")

		// 		for _, loc := range locales {

		// 			if loc == baseLocale {

		// 				// plural rule found
		// 				fmt.Println("Locale Plural Rules Found:", loc, baseLocale, l)
		// 			}
		// 		}
		// 		// fmt.Println(pr.Locales)

		// 		// for _, rule := range pr.PluralRule {
		// 		// 	fmt.Println(rule.Count, rule.Common.Data())
		// 		// }
		// 	}
		// }

		// End Plural Rules

		if err = os.MkdirAll(fmt.Sprintf(locDir, l), 0777); err != nil {
			log.Fatal(err)
		}

		filename := fmt.Sprintf(locFilename, l, l)

		output, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer output.Close()

		if err := tmpl.ExecuteTemplate(output, "translator", trans); err != nil {
			log.Fatal(err)
		}

		output.Close()

		// after file written run gofmt on file to ensure best formatting
		cmd := exec.Command("gofmt", "-s", "-w", filename)
		if err = cmd.Run(); err != nil {
			log.Panic(err)
		}
	}
}

// TODO: cleanup function logic perhaps write a lexer... but it's working right now, and
// I'm already farther down the rabbit hole than I'd like and so pulling the chute here.
func parseCardinalPluralRuleFunc(current *cldr.CLDR, baseLocale string) (results string, plurals string) {

	var prCardinal *struct {
		cldr.Common
		Locales    string "xml:\"locales,attr\""
		PluralRule []*struct {
			cldr.Common
			Count string "xml:\"count,attr\""
		} "xml:\"pluralRule\""
	}

	var pluralArr []locales.PluralRule

	for _, p := range current.Supplemental().Plurals {

		for _, pr := range p.PluralRules {

			locs := strings.Split(pr.Locales, " ")

			for _, loc := range locs {

				if loc == baseLocale {
					prCardinal = pr
				}
			}
		}
	}

	// no plural rules for locale
	if prCardinal == nil {
		plurals = "nil"
		results = "return locales.PluralRuleUnknown,nil"
		return
	}

	vals := make(map[string]struct{})
	first := true

	// pre parse for variables
	for _, rule := range prCardinal.PluralRule {

		ps1 := pluralStringToString(rule.Count)
		psI := pluralStringToInt(rule.Count)
		pluralArr = append(pluralArr, psI)

		// fmt.Println(rule.Count, ps1)

		data := strings.Replace(strings.Replace(strings.Replace(strings.TrimSpace(strings.SplitN(rule.Common.Data(), "@", 2)[0]), " = ", " == ", -1), " or ", " || ", -1), " and ", " && ", -1)

		if len(data) == 0 {
			if len(prCardinal.PluralRule) == 1 {

				results = "return locales." + ps1 + ", nil"

			} else {

				results += "\n\nreturn locales." + ps1 + ", nil"
				// results += "else {\nreturn locales." + locales.PluralStringToString(rule.Count) + ", nil\n}"
			}

			continue
		}

		if strings.Contains(data, "n") {
			vals[prVarFuncs["n"]] = struct{}{}
		}

		if strings.Contains(data, "i") {
			vals[prVarFuncs["i"]] = struct{}{}
		}

		if strings.Contains(data, "v") {
			vals[prVarFuncs["v"]] = struct{}{}
		}

		if strings.Contains(data, "w") {
			vals[prVarFuncs["w"]] = struct{}{}
		}

		if strings.Contains(data, "f") {
			vals[prVarFuncs["f"]] = struct{}{}
		}

		if strings.Contains(data, "t") {
			vals[prVarFuncs["t"]] = struct{}{}
		}

		if first {
			results += "if "
			first = false
		} else {
			results += "else if "
		}

		stmt := ""

		// real work here
		//
		// split by 'or' then by 'and' allowing to better
		// determine bracketing for formula

		ors := strings.Split(data, "||")

		for _, or := range ors {

			stmt += "("

			ands := strings.Split(strings.TrimSpace(or), "&&")

			for _, and := range ands {

				inArg := false
				pre := ""
				lft := ""
				preOperator := ""
				args := strings.Split(strings.TrimSpace(and), " ")

				for _, a := range args {

					if inArg {
						// check to see if is a value range 2..9

						multiRange := strings.Count(a, "..") > 1
						cargs := strings.Split(strings.TrimSpace(a), ",")
						hasBracket := len(cargs) > 1
						bracketAdded := false
						lastWasRange := false

						for _, carg := range cargs {

							if rng := strings.Split(carg, ".."); len(rng) > 1 {

								if multiRange {
									pre += " ("
								} else {
									pre += " "
								}

								switch preOperator {
								case "==":
									pre += lft + " >= " + rng[0] + " && " + lft + "<=" + rng[1]
								case "!=":
									pre += lft + " < " + rng[0] + " && " + lft + " > " + rng[1]
								}

								if multiRange {
									pre += ") || "
								} else {
									pre += " || "
								}

								lastWasRange = true
								continue
							}

							if lastWasRange {
								pre = strings.TrimRight(pre, " || ") + " && "
							}

							lastWasRange = false

							if hasBracket && !bracketAdded {
								pre += "("
								bracketAdded = true
							}

							// single comma separated values
							switch preOperator {
							case "==":
								pre += " " + lft + preOperator + carg + " || "
							case "!=":
								pre += " " + lft + preOperator + carg + " && "
							}

						}

						pre = strings.TrimRight(pre, " || ")
						pre = strings.TrimRight(pre, " && ")
						pre = strings.TrimRight(pre, " || ")

						if hasBracket && bracketAdded {
							pre += ")"
						}

						continue
					}

					if strings.Contains(a, "=") || a == ">" || a == "<" {
						inArg = true
						preOperator = a
						continue
					}

					lft += a
				}

				stmt += pre + " && "
			}

			stmt = strings.TrimRight(stmt, " && ") + ") || "
		}

		stmt = strings.TrimRight(stmt, " || ")

		results += stmt

		results += " {\n"

		// return plural rule here
		results += "return locales." + ps1 + ", nil\n"

		results += "}"
	}

	pre := "\n"

	for k := range vals {
		pre += k
	}

	pre += "\n"

	if len(results) == 0 {
		results = "return locales.PluralRuleUnknown,nil"
	}

	results = pre + results

	if len(pluralArr) == 0 {
		plurals = "nil"
	} else {
		plurals = fmt.Sprintf("%#v", pluralArr)
	}

	return
}

// pluralStringToInt returns the enum value of 'plural' provided
func pluralStringToInt(plural string) locales.PluralRule {

	switch plural {
	case "zero":
		return locales.PluralRuleZero
	case "one":
		return locales.PluralRuleOne
	case "two":
		return locales.PluralRuleTwo
	case "few":
		return locales.PluralRuleFew
	case "many":
		return locales.PluralRuleMany
	case "other":
		return locales.PluralRuleOther
	default:
		return locales.PluralRuleUnknown
	}
}

func pluralStringToString(pr string) string {

	pr = strings.TrimSpace(pr)

	switch pr {
	case "zero":
		return "PluralRuleZero"
	case "one":
		return "PluralRuleOne"
	case "two":
		return "PluralRuleTwo"
	case "few":
		return "PluralRuleFew"
	case "many":
		return "PluralRuleMany"
	case "other":
		return "PluralRuleOther"
	default:
		return "PluralRuleUnknown"
	}
}

// //plurals
// rules := "data/rules"
// plurals := map[string]*pluralInfo{}
// basePlurals := map[string]string{}

// err := filepath.Walk(rules, func(path string, info os.FileInfo, err error) error {

// 	if err != nil {
// 		panic(err)
// 	}

// 	if info.IsDir() {
// 		return nil
// 	}

// 	in, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var out yaml.MapSlice
// 	if err = yaml.Unmarshal(in, &out); err != nil {
// 		panic(err)
// 	}

// 	var plural string
// 	for _, item := range out {
// 		if item.Key == "plural" {
// 			plural = item.Value.(string)
// 		}
// 	}

// 	locale := strings.Replace(info.Name(), filepath.Ext(info.Name()), "", 1)
// 	locale = strings.ToLower(strings.Replace(locale, "-", "_", -1))

// 	plurals[locale] = &pluralInfo{
// 		path:   path,
// 		locale: locale,
// 		plural: plural,
// 	}

// 	if plural == "" {
// 		return nil
// 	}

// 	basePlurals[locale] = plural

// 	return nil
// })

// if err != nil {
// 	panic(err)
// }

// for _, p := range plurals {

// 	if p.plural == "" {

// 		var ok bool

// 		fmt.Print("can't find plurals in ", p.path, " attempting to locate base language plural rules...")

// 		base := strings.SplitN(p.locale, "_", 2)

// 		p.plural, ok = basePlurals[base[0]]
// 		if !ok {
// 			fmt.Println("Not Found")
// 			continue
// 		}

// 		fmt.Println("Found")
// 	}
// }

// cldr

// var decoder cldr.Decoder
// cldr, err := decoder.DecodePath("data/core")
// if err != nil {
// 	panic(err)
// }

// 	locs := map[string]string{}
// 	numbers := map[string]i18n.Number{}
// 	calendars := map[string]i18n.Calendar{}
// 	locales := cldr.Locales()
// 	for _, loc := range locales {

// 		ldml := cldr.RawLDML(loc)
// 		if ldml.Numbers == nil {
// 			continue
// 		}
// 		var number i18n.Number

// 		number.Currencies = make(i18n.CurrencyFormatValue)

// 		if len(ldml.Numbers.Symbols) > 0 {
// 			symbol := ldml.Numbers.Symbols[0]
// 			if len(symbol.Decimal) > 0 {
// 				number.Symbols.Decimal = symbol.Decimal[0].Data()
// 			}
// 			if len(symbol.Group) > 0 {
// 				number.Symbols.Group = symbol.Group[0].Data()
// 			}
// 			if len(symbol.MinusSign) > 0 {
// 				number.Symbols.Negative = symbol.MinusSign[0].Data()
// 			}
// 			if len(symbol.PercentSign) > 0 {
// 				number.Symbols.Percent = symbol.PercentSign[0].Data()
// 			}
// 			if len(symbol.PerMille) > 0 {
// 				number.Symbols.PerMille = symbol.PerMille[0].Data()
// 			}
// 		}
// 		if len(ldml.Numbers.DecimalFormats) > 0 && len(ldml.Numbers.DecimalFormats[0].DecimalFormatLength) > 0 {
// 			number.Formats.Decimal = ldml.Numbers.DecimalFormats[0].DecimalFormatLength[0].DecimalFormat[0].Pattern[0].Data()
// 		}

// 		if len(ldml.Numbers.CurrencyFormats) > 0 && len(ldml.Numbers.CurrencyFormats[0].CurrencyFormatLength) > 0 {

// 			number.Formats.Currency = ldml.Numbers.CurrencyFormats[0].CurrencyFormatLength[0].CurrencyFormat[0].Pattern[0].Data()
// 			number.Formats.CurrencyAccounting = number.Formats.Currency

// 			if len(ldml.Numbers.CurrencyFormats[0].CurrencyFormatLength[0].CurrencyFormat) > 1 {
// 				number.Formats.CurrencyAccounting = ldml.Numbers.CurrencyFormats[0].CurrencyFormatLength[0].CurrencyFormat[1].Pattern[0].Data()
// 			}
// 		}
// 		if len(ldml.Numbers.PercentFormats) > 0 && len(ldml.Numbers.PercentFormats[0].PercentFormatLength) > 0 {
// 			number.Formats.Percent = ldml.Numbers.PercentFormats[0].PercentFormatLength[0].PercentFormat[0].Pattern[0].Data()
// 		}
// 		if ldml.Numbers.Currencies != nil {

// 			for _, currency := range ldml.Numbers.Currencies.Currency {

// 				var c i18n.Currency

// 				c.Currency = currency.Type

// 				if len(currency.DisplayName) > 0 {
// 					c.DisplayName = currency.DisplayName[0].Data()
// 				}

// 				if len(currency.Symbol) > 0 {
// 					c.Symbol = currency.Symbol[0].Data()
// 				}

// 				number.Currencies[c.Currency] = c
// 			}
// 		}
// 		numbers[loc] = number
// 		locs[loc] = strings.ToLower(strings.Replace(loc, "_", "", -1))

// 		if ldml.Dates != nil && ldml.Dates.Calendars != nil {

// 			var calendar i18n.Calendar
// 			ldmlCar := ldml.Dates.Calendars.Calendar[0]

// 			for _, cal := range ldml.Dates.Calendars.Calendar {
// 				if cal.Type == "gregorian" {
// 					ldmlCar = cal
// 				}
// 			}

// 			if ldmlCar.DateFormats != nil {
// 				for _, datefmt := range ldmlCar.DateFormats.DateFormatLength {
// 					switch datefmt.Type {
// 					case "full":
// 						calendar.Formats.DateEra.BC.Full = datefmt.DateFormat[0].Pattern[0].Data()
// 						calendar.Formats.DateEra.AD.Full = calendar.Formats.DateEra.BC.Full

// 					case "long":
// 						calendar.Formats.DateEra.BC.Long = datefmt.DateFormat[0].Pattern[0].Data()
// 						calendar.Formats.DateEra.AD.Long = calendar.Formats.DateEra.BC.Long

// 						// Overrides TODO: Split Each Section into separate function, getting to big to maintain

// 					case "medium":
// 						calendar.Formats.DateEra.BC.Medium = datefmt.DateFormat[0].Pattern[0].Data()
// 						calendar.Formats.DateEra.AD.Medium = calendar.Formats.DateEra.BC.Medium
// 					case "short":
// 						calendar.Formats.DateEra.BC.Short = datefmt.DateFormat[0].Pattern[0].Data()
// 						calendar.Formats.DateEra.AD.Short = calendar.Formats.DateEra.BC.Short
// 					}
// 				}

// 				// Overrides TODO: Split Each Section into separate function, getting to big to maintain
// 				switch loc {
// 				case "th":
// 					calendar.Formats.DateEra.BC.Full = "EEEEที่ d MMMM y GGGG"
// 					calendar.Formats.DateEra.AD.Full = "EEEEที่ d MMMM GGGG y"
// 					calendar.Formats.DateEra.BC.Long = "d MMMM y GG"
// 					calendar.Formats.DateEra.AD.Long = "d MMMM GG y"
// 				case "en":
// 					calendar.Formats.DateEra.BC.Full = "EEEE, MMMM d, y GGGG"
// 					calendar.Formats.DateEra.BC.Long = "MMMM d, y GG"
// 					// calendar.Formats.DateEra.BC.Medium = "MMM d, y GG"
// 					// calendar.Formats.DateEra.BC.Short = "M/d/yy G"
// 				}
// 			}

// 			if ldmlCar.TimeFormats != nil {
// 				for _, datefmt := range ldmlCar.TimeFormats.TimeFormatLength {
// 					switch datefmt.Type {
// 					case "full":
// 						calendar.Formats.Time.Full = datefmt.TimeFormat[0].Pattern[0].Data()
// 					case "long":
// 						calendar.Formats.Time.Long = datefmt.TimeFormat[0].Pattern[0].Data()
// 					case "medium":
// 						calendar.Formats.Time.Medium = datefmt.TimeFormat[0].Pattern[0].Data()
// 					case "short":
// 						calendar.Formats.Time.Short = datefmt.TimeFormat[0].Pattern[0].Data()
// 					}
// 				}
// 			}
// 			if ldmlCar.DateTimeFormats != nil {
// 				for _, datefmt := range ldmlCar.DateTimeFormats.DateTimeFormatLength {
// 					switch datefmt.Type {
// 					case "full":
// 						calendar.Formats.DateTime.Full = datefmt.DateTimeFormat[0].Pattern[0].Data()
// 					case "long":
// 						calendar.Formats.DateTime.Long = datefmt.DateTimeFormat[0].Pattern[0].Data()
// 					case "medium":
// 						calendar.Formats.DateTime.Medium = datefmt.DateTimeFormat[0].Pattern[0].Data()
// 					case "short":
// 						calendar.Formats.DateTime.Short = datefmt.DateTimeFormat[0].Pattern[0].Data()
// 					}
// 				}
// 			}
// 			if ldmlCar.Months != nil {

// 				for _, monthctx := range ldmlCar.Months.MonthContext {

// 					for _, months := range monthctx.MonthWidth {

// 						i18nMonth := make(i18n.CalendarMonthFormatNameValue)

// 						for _, m := range months.Month {
// 							switch m.Type {
// 							case "1":
// 								i18nMonth[time.January] = m.Data()
// 							case "2":
// 								i18nMonth[time.February] = m.Data()
// 							case "3":
// 								i18nMonth[time.March] = m.Data()
// 							case "4":
// 								i18nMonth[time.April] = m.Data()
// 							case "5":
// 								i18nMonth[time.May] = m.Data()
// 							case "6":
// 								i18nMonth[time.June] = m.Data()
// 							case "7":
// 								i18nMonth[time.July] = m.Data()
// 							case "8":
// 								i18nMonth[time.August] = m.Data()
// 							case "9":
// 								i18nMonth[time.September] = m.Data()
// 							case "10":
// 								i18nMonth[time.October] = m.Data()
// 							case "11":
// 								i18nMonth[time.November] = m.Data()
// 							case "12":
// 								i18nMonth[time.December] = m.Data()
// 							}
// 						}
// 						switch months.Type {
// 						case "abbreviated":
// 							calendar.FormatNames.Months.Abbreviated = i18nMonth
// 						case "narrow":
// 							calendar.FormatNames.Months.Narrow = i18nMonth
// 						case "short":
// 							calendar.FormatNames.Months.Short = i18nMonth
// 						case "wide":
// 							calendar.FormatNames.Months.Wide = i18nMonth
// 						}
// 					}
// 				}
// 			}
// 			if ldmlCar.Days != nil {
// 				for _, dayctx := range ldmlCar.Days.DayContext {

// 					for _, days := range dayctx.DayWidth {

// 						i18nDay := make(i18n.CalendarDayFormatNameValue)

// 						for _, d := range days.Day {

// 							switch d.Type {
// 							case "sun":
// 								i18nDay[time.Sunday] = d.Data()
// 							case "mon":
// 								i18nDay[time.Monday] = d.Data()
// 							case "tue":
// 								i18nDay[time.Tuesday] = d.Data()
// 							case "wed":
// 								i18nDay[time.Wednesday] = d.Data()
// 							case "thu":
// 								i18nDay[time.Thursday] = d.Data()
// 							case "fri":
// 								i18nDay[time.Friday] = d.Data()
// 							case "sat":
// 								i18nDay[time.Saturday] = d.Data()
// 							}
// 						}

// 						switch days.Type {
// 						case "abbreviated":
// 							calendar.FormatNames.Days.Abbreviated = i18nDay
// 						case "narrow":
// 							calendar.FormatNames.Days.Narrow = i18nDay
// 						case "short":
// 							calendar.FormatNames.Days.Short = i18nDay
// 						case "wide":
// 							calendar.FormatNames.Days.Wide = i18nDay
// 						}
// 					}
// 				}
// 			}

// 			if ldmlCar.DayPeriods != nil {

// 				for _, ctx := range ldmlCar.DayPeriods.DayPeriodContext {

// 					for _, width := range ctx.DayPeriodWidth {

// 						// var i18nPeriod i18n.CalendarPeriodFormatNameValue
// 						i18nPeriod := make(i18n.CalendarPeriodFormatNameValue)

// 						for _, d := range width.DayPeriod {

// 							if _, ok := i18nPeriod[d.Type]; !ok {
// 								i18nPeriod[d.Type] = d.Data()
// 							}
// 						}

// 						switch width.Type {
// 						case "abbreviated":
// 							calendar.FormatNames.Periods.Abbreviated = i18nPeriod
// 						case "narrow":
// 							calendar.FormatNames.Periods.Narrow = i18nPeriod
// 						case "short":
// 							calendar.FormatNames.Periods.Short = i18nPeriod
// 						case "wide":
// 							calendar.FormatNames.Periods.Wide = i18nPeriod
// 						}
// 					}
// 				}
// 				// var empty i18n.CalendarPeriodFormatNameValue
// 				// if calendar.FormatNames.Periods.Abbreviated == empty {
// 				// 	calendar.FormatNames.Periods.Abbreviated = calendar.FormatNames.Periods.Wide
// 				// }
// 			}

// 			if ldmlCar.Eras != nil {

// 				var eras i18n.Eras

// 				if ldmlCar.Eras.EraNames != nil && len(ldmlCar.Eras.EraNames.Era) > 0 {
// 					eras.BC.Full = ldmlCar.Eras.EraNames.Era[0].Data()
// 				}

// 				if ldmlCar.Eras.EraAbbr != nil && len(ldmlCar.Eras.EraAbbr.Era) > 0 {
// 					eras.BC.Abbrev = ldmlCar.Eras.EraAbbr.Era[0].Data()
// 				}

// 				if ldmlCar.Eras.EraNarrow != nil && len(ldmlCar.Eras.EraNarrow.Era) > 0 {
// 					eras.BC.Narrow = ldmlCar.Eras.EraNarrow.Era[0].Data()
// 				}

// 				if ldmlCar.Eras.EraNames != nil && len(ldmlCar.Eras.EraNames.Era) > 2 {
// 					eras.AD.Full = ldmlCar.Eras.EraNames.Era[2].Data()
// 				}

// 				if ldmlCar.Eras.EraAbbr != nil && len(ldmlCar.Eras.EraAbbr.Era) > 2 {
// 					eras.AD.Abbrev = ldmlCar.Eras.EraAbbr.Era[2].Data()
// 				}

// 				if ldmlCar.Eras.EraNarrow != nil && len(ldmlCar.Eras.EraNarrow.Era) > 2 {
// 					eras.AD.Narrow = ldmlCar.Eras.EraNarrow.Era[2].Data()
// 				}

// 				calendar.FormatNames.Eras = eras
// 			}

// 			calendars[loc] = calendar
// 		}
// 	}

// 	for locale := range locs {

// 		if !strings.Contains(locale, "_") {
// 			continue
// 		}

// 		calendar := calendars[locale]

// 		bString := strings.SplitN(locale, "_", 2)
// 		base := bString[0]

// 		baseCal := calendars[base]

// 		// copy base calendar objects

// 		// Dates
// 		if calendar.Formats.DateEra.AD.Full == "" {
// 			calendar.Formats.DateEra.BC.Full = baseCal.Formats.DateEra.BC.Full
// 			calendar.Formats.DateEra.AD.Full = baseCal.Formats.DateEra.AD.Full
// 		}

// 		if calendar.Formats.DateEra.AD.Long == "" {
// 			calendar.Formats.DateEra.BC.Long = baseCal.Formats.DateEra.BC.Long
// 			calendar.Formats.DateEra.AD.Long = baseCal.Formats.DateEra.AD.Long
// 		}

// 		if calendar.Formats.DateEra.AD.Medium == "" {
// 			calendar.Formats.DateEra.BC.Medium = baseCal.Formats.DateEra.BC.Medium
// 			calendar.Formats.DateEra.AD.Medium = baseCal.Formats.DateEra.AD.Medium
// 		}

// 		if calendar.Formats.DateEra.AD.Short == "" {
// 			calendar.Formats.DateEra.BC.Short = baseCal.Formats.DateEra.BC.Short
// 			calendar.Formats.DateEra.AD.Short = baseCal.Formats.DateEra.AD.Short
// 		}

// 		// times
// 		if calendar.Formats.Time.Full == "" {
// 			calendar.Formats.Time.Full = baseCal.Formats.Time.Full
// 		}

// 		if calendar.Formats.Time.Long == "" {
// 			calendar.Formats.Time.Long = baseCal.Formats.Time.Long
// 		}

// 		if calendar.Formats.Time.Medium == "" {
// 			calendar.Formats.Time.Medium = baseCal.Formats.Time.Medium
// 		}

// 		if calendar.Formats.Time.Short == "" {
// 			calendar.Formats.Time.Short = baseCal.Formats.Time.Short
// 		}

// 		// date & times
// 		if calendar.Formats.DateTime.Full == "" {
// 			calendar.Formats.DateTime.Full = baseCal.Formats.DateTime.Full
// 		}

// 		if calendar.Formats.DateTime.Long == "" {
// 			calendar.Formats.DateTime.Long = baseCal.Formats.DateTime.Long
// 		}

// 		if calendar.Formats.DateTime.Medium == "" {
// 			calendar.Formats.DateTime.Medium = baseCal.Formats.DateTime.Medium
// 		}

// 		if calendar.Formats.DateTime.Short == "" {
// 			calendar.Formats.DateTime.Short = baseCal.Formats.DateTime.Short
// 		}

// 		// months

// 		if calendar.FormatNames.Months.Abbreviated == nil {
// 			calendar.FormatNames.Months.Abbreviated = make(i18n.CalendarMonthFormatNameValue)
// 		}

// 		if calendar.FormatNames.Months.Narrow == nil {
// 			calendar.FormatNames.Months.Narrow = make(i18n.CalendarMonthFormatNameValue)
// 		}

// 		if calendar.FormatNames.Months.Short == nil {
// 			calendar.FormatNames.Months.Short = make(i18n.CalendarMonthFormatNameValue)
// 		}

// 		if calendar.FormatNames.Months.Wide == nil {
// 			calendar.FormatNames.Months.Wide = make(i18n.CalendarMonthFormatNameValue)
// 		}

// 		for k, v := range baseCal.FormatNames.Months.Abbreviated {

// 			val, ok := calendar.FormatNames.Months.Abbreviated[k]
// 			if !ok {
// 				calendar.FormatNames.Months.Abbreviated[k] = v
// 				continue
// 			}

// 			if val == "" {
// 				calendar.FormatNames.Months.Abbreviated[k] = v
// 			}
// 		}

// 		for k, v := range baseCal.FormatNames.Months.Narrow {

// 			val, ok := calendar.FormatNames.Months.Narrow[k]
// 			if !ok {
// 				calendar.FormatNames.Months.Narrow[k] = v
// 				continue
// 			}

// 			if val == "" {
// 				calendar.FormatNames.Months.Narrow[k] = v
// 			}
// 		}

// 		for k, v := range baseCal.FormatNames.Months.Short {

// 			val, ok := calendar.FormatNames.Months.Short[k]
// 			if !ok {
// 				calendar.FormatNames.Months.Short[k] = v
// 				continue
// 			}

// 			if val == "" {
// 				calendar.FormatNames.Months.Short[k] = v
// 			}
// 		}

// 		for k, v := range baseCal.FormatNames.Months.Wide {

// 			val, ok := calendar.FormatNames.Months.Wide[k]
// 			if !ok {
// 				calendar.FormatNames.Months.Wide[k] = v
// 				continue
// 			}

// 			if val == "" {
// 				calendar.FormatNames.Months.Wide[k] = v
// 			}
// 		}

// 		// days

// 		if calendar.FormatNames.Days.Abbreviated == nil {
// 			calendar.FormatNames.Days.Abbreviated = make(i18n.CalendarDayFormatNameValue)
// 		}

// 		if calendar.FormatNames.Days.Narrow == nil {
// 			calendar.FormatNames.Days.Narrow = make(i18n.CalendarDayFormatNameValue)
// 		}

// 		if calendar.FormatNames.Days.Short == nil {
// 			calendar.FormatNames.Days.Short = make(i18n.CalendarDayFormatNameValue)
// 		}

// 		if calendar.FormatNames.Days.Wide == nil {
// 			calendar.FormatNames.Days.Wide = make(i18n.CalendarDayFormatNameValue)
// 		}

// 		for k, v := range baseCal.FormatNames.Days.Abbreviated {

// 			val, ok := calendar.FormatNames.Days.Abbreviated[k]
// 			if !ok {
// 				calendar.FormatNames.Days.Abbreviated[k] = v
// 				continue
// 			}

// 			if val == "" {
// 				calendar.FormatNames.Days.Abbreviated[k] = v
// 			}
// 		}

// 		for k, v := range baseCal.FormatNames.Days.Narrow {

// 			val, ok := calendar.FormatNames.Days.Narrow[k]
// 			if !ok {
// 				calendar.FormatNames.Days.Narrow[k] = v
// 				continue
// 			}

// 			if val == "" {
// 				calendar.FormatNames.Days.Narrow[k] = v
// 			}
// 		}

// 		for k, v := range baseCal.FormatNames.Days.Short {

// 			val, ok := calendar.FormatNames.Days.Short[k]
// 			if !ok {
// 				calendar.FormatNames.Days.Short[k] = v
// 				continue
// 			}

// 			if val == "" {
// 				calendar.FormatNames.Days.Short[k] = v
// 			}
// 		}

// 		for k, v := range baseCal.FormatNames.Days.Wide {

// 			val, ok := calendar.FormatNames.Days.Wide[k]
// 			if !ok {
// 				calendar.FormatNames.Days.Wide[k] = v
// 				continue
// 			}

// 			if val == "" {
// 				calendar.FormatNames.Days.Wide[k] = v
// 			}
// 		}

// 		// periods
// 		if calendar.FormatNames.Periods.Abbreviated == nil {
// 			calendar.FormatNames.Periods.Abbreviated = make(i18n.CalendarPeriodFormatNameValue)
// 		}

// 		if calendar.FormatNames.Periods.Narrow == nil {
// 			calendar.FormatNames.Periods.Narrow = make(i18n.CalendarPeriodFormatNameValue)
// 		}

// 		if calendar.FormatNames.Periods.Short == nil {
// 			calendar.FormatNames.Periods.Short = make(i18n.CalendarPeriodFormatNameValue)
// 		}

// 		if calendar.FormatNames.Periods.Wide == nil {
// 			calendar.FormatNames.Periods.Wide = make(i18n.CalendarPeriodFormatNameValue)
// 		}

// 		for k, v := range baseCal.FormatNames.Periods.Abbreviated {

// 			val, ok := calendar.FormatNames.Periods.Abbreviated[k]
// 			if !ok {
// 				calendar.FormatNames.Periods.Abbreviated[k] = v
// 				continue
// 			}

// 			if val == "" {
// 				calendar.FormatNames.Periods.Abbreviated[k] = v
// 			}
// 		}

// 		for k, v := range baseCal.FormatNames.Periods.Narrow {

// 			val, ok := calendar.FormatNames.Periods.Narrow[k]
// 			if !ok {
// 				calendar.FormatNames.Periods.Narrow[k] = v
// 				continue
// 			}

// 			if val == "" {
// 				calendar.FormatNames.Periods.Narrow[k] = v
// 			}
// 		}

// 		for k, v := range baseCal.FormatNames.Periods.Short {

// 			val, ok := calendar.FormatNames.Periods.Short[k]
// 			if !ok {
// 				calendar.FormatNames.Periods.Short[k] = v
// 				continue
// 			}

// 			if val == "" {
// 				calendar.FormatNames.Periods.Short[k] = v
// 			}
// 		}

// 		for k, v := range baseCal.FormatNames.Periods.Wide {

// 			val, ok := calendar.FormatNames.Periods.Wide[k]
// 			if !ok {
// 				calendar.FormatNames.Periods.Wide[k] = v
// 				continue
// 			}

// 			if val == "" {
// 				calendar.FormatNames.Periods.Wide[k] = v
// 			}
// 		}

// 		calendars[locale] = calendar

// 		number := numbers[locale]
// 		baseNum := numbers[base]

// 		// symbols
// 		if number.Symbols.Decimal == "" {
// 			number.Symbols.Decimal = baseNum.Symbols.Decimal
// 		}

// 		if number.Symbols.Group == "" {
// 			number.Symbols.Group = baseNum.Symbols.Group
// 		}

// 		if number.Symbols.Negative == "" {
// 			number.Symbols.Negative = baseNum.Symbols.Negative
// 		}

// 		if number.Symbols.Percent == "" {
// 			number.Symbols.Percent = baseNum.Symbols.Percent
// 		}

// 		if number.Symbols.PerMille == "" {
// 			number.Symbols.PerMille = baseNum.Symbols.PerMille
// 		}

// 		// formats
// 		if number.Formats.Decimal == "" {
// 			number.Formats.Decimal = baseNum.Formats.Decimal
// 		}

// 		if number.Formats.Currency == "" {
// 			number.Formats.Currency = baseNum.Formats.Currency
// 		}

// 		if number.Formats.CurrencyAccounting == "" {
// 			number.Formats.CurrencyAccounting = baseNum.Formats.CurrencyAccounting
// 		}

// 		if number.Formats.Percent == "" {
// 			number.Formats.Percent = baseNum.Formats.Percent
// 		}

// 		// currency
// 		for k, v := range baseNum.Currencies {

// 			val, ok := number.Currencies[k]
// 			if !ok {
// 				// number.Currencies[k] = v
// 				continue
// 			}

// 			if val.Currency == "" {
// 				val.Currency = v.Currency
// 			}

// 			if val.DisplayName == "" {
// 				val.DisplayName = v.DisplayName
// 			}

// 			if val.Symbol == "" {
// 				val.Symbol = v.Symbol
// 			}

// 			number.Currencies[k] = val
// 		}

// 		numbers[locale] = number
// 	}

// 	var wg sync.WaitGroup
// 	wg.Add(len(numbers))
// 	for locale, number := range numbers {
// 		go func(locale string, number i18n.Number) {

// 			localeLowercase := strings.ToLower(locale)

// 			defer func() { wg.Done() }()
// 			path := "../../resources/locales/" + locale

// 			if _, err := os.Stat(path); err != nil {
// 				if err = os.MkdirAll(path, 0777); err != nil {
// 					panic(err)
// 				}
// 			}

// 			path += "/"

// 			mainFile, err := os.Create(path + "main.go")
// 			if err != nil {
// 				panic(err)
// 			}
// 			defer mainFile.Close()

// 			calendar := calendars[locale]

// 			mainCodes, err := format.Source([]byte(fmt.Sprintf(`package %s
// 			import "github.com/go-playground/universal-translator"

// 			var locale = &ut.Locale{
// 				Locale: %q,
// 				Number: ut.Number{
// 					Symbols: symbols,
// 					Formats: formats,
// 					Currencies: currencies,
// 				},
// 				Calendar: calendar,
// 				PluralRule:   pluralRule,
// 			}

// 			func init() {
// 				ut.RegisterLocale(locale)
// 			}
// 		`, locale, locale)))

// 			if err != nil {
// 				panic(err)
// 			}

// 			fmt.Fprintf(mainFile, "%s", mainCodes)

// 			numberFile, err := os.Create(path + "number.go")
// 			if err != nil {
// 				panic(err)
// 			}
// 			defer numberFile.Close()

// 			numberCodes, err := format.Source([]byte(fmt.Sprintf(`package %s
// 			import "github.com/go-playground/universal-translator"
// 			var (
// 				symbols = %#v
// 				formats = %#v
// 			)
// 		`, locale, number.Symbols, number.Formats)))
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Fprintf(numberFile, "%s", numberCodes)

// 			currencyFile, err := os.Create(path + "currency.go")
// 			if err != nil {
// 				panic(err)
// 			}
// 			defer currencyFile.Close()

// 			currencyCodes, err := format.Source([]byte(fmt.Sprintf(`package %s
// 			import "github.com/go-playground/universal-translator"
// 			var currencies = %# v
// 		`, locale, number.Currencies)))
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Fprintf(currencyFile, "%s", currencyCodes)

// 			calendarFile, err := os.Create(path + "calendar.go")
// 			if err != nil {
// 				panic(err)
// 			}
// 			defer calendarFile.Close()

// 			calendarCodes, err := format.Source([]byte(fmt.Sprintf(`package %s
// 			import "github.com/go-playground/universal-translator"
// 			var calendar = %#v
// 		`, locale, calendar)))
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Fprintf(calendarFile, "%s", calendarCodes)

// 			var ok bool
// 			pluralCode := "1"

// 			pInfo, ok := plurals[localeLowercase]
// 			if ok && pInfo.plural != "" {
// 				pluralCode = pInfo.plural
// 			}

// 			pluralFile, err := os.Create(path + "plural.go")
// 			if err != nil {
// 				panic(err)
// 			}
// 			defer pluralFile.Close()

// 			pluralCodes, err := format.Source([]byte(fmt.Sprintf(`package %s

// 			var pluralRule = %q
// 		`, locale, pluralCode)))
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Fprintf(pluralFile, "%s", pluralCodes)
// 		}(locale, number)
// 	}

// 	wg.Wait()

// 	localesFile, err := os.Create("../../resources/locales/all.go")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer localesFile.Close()

// 	tmpl, err := template.New("").Parse(`package locales

// 		// Imports for all locales
// 		import (
// 			{{range $locale, $_ := .}}// Locale "{{$locale}}" import that automatically registers itslef with the universal-translator package
// 			_ "github.com/go-playground/universal-translator/resources/locales/{{$locale}}"
// 		{{end}})
// 	`)

// 	if err != nil {
// 		panic(err)
// 	}
// 	var buf bytes.Buffer
// 	if err := tmpl.Execute(&buf, locs); err != nil {
// 		panic(err)
// 	}
// 	allCodes, err := format.Source(buf.Bytes())
// 	if err != nil {
// 		panic(err)
// 	}
// 	_, err = localesFile.Write(allCodes)
// 	if err != nil {
// 		panic(err)
// 	}
// }
