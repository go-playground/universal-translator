package main

import (
	"fmt"
	"time"

	"github.com/go-playground/universal-translator"

	// DONE this way to avoid huge compile times + memory for all languages, although it would
	// be nice for all applications to support all languages... that's not reality
	_ "github.com/go-playground/universal-translator/resources/locales"
)

func main() {
	trans, _ := ut.GetTranslator("en")

	trans.PrintPluralRules()
	// OUTPUT:
	// Translator locale 'en' supported rules:
	//- PluralRuleOne
	//- PluralRuleOther

	// add a singular translation
	trans.Add(ut.PluralRuleOne, "homepage", "welcome_msg", "Welcome to site %s")

	// add singular + plural translation(s)
	trans.Add(ut.PluralRuleOne, "homepage", "day_warning", "You only have %d day left in your trial")
	trans.Add(ut.PluralRuleOther, "homepage", "day_warning", "You only have %d day's left in your trial")

	// translate singular
	translated := trans.T("welcome_msg", "Joey Bloggs")
	fmt.Println(translated)
	// OUTPUT: Welcome to site Joey Bloggs

	// What if something went wrong? then translated would output "" (blank)
	// How do I catch errors?
	translated, err := trans.TSafe("welcome_m", "Joey Bloggs")
	fmt.Println(translated)
	// OUTPUT: ""
	fmt.Println(err)
	// OUTPUT: ***** WARNING:***** Translation Key 'welcome_m' Not Found

	// NOTE: there is a Safe variant of most of the Translation and Formatting functions if you need them,
	// for brevity will be using the non safe ones for the rest of this example

	// The second parameter below, count, is needed as the final variable is a varadic and would not
	// know which one to use in applying the plural rules.
	// translate singular/plural
	translated = trans.P("day_warning", 3, 3)
	fmt.Println(translated)
	// OUTPUT: You only have 3 day's left in your trial

	translated = trans.P("day_warning", 1, 1)
	fmt.Println(translated)
	// OUTPUT: You only have 1 day left in your trial

	// There are Full, Long, Medium and Short function for each of the following
	dtString := "Jan 2, 2006 at 3:04:05pm"
	dt, _ := time.Parse(dtString, dtString)

	formatted := trans.FmtDateFull(dt)
	fmt.Println(formatted)
	// OUTPUT: Monday, January 2, 2006

	formatted = trans.FmtDateShort(dt)
	fmt.Println(formatted)
	// OUTPUT: 1/2/06

	formatted = trans.FmtTimeFull(dt)
	fmt.Println(formatted)
	// OUTPUT: 3:04:05 PM

	formatted = trans.FmtDateTimeFull(dt)
	fmt.Println(formatted)
	// OUTPUT: Monday, January 2, 2006 at 3:04:05 PM

	formatted = trans.FmtCurrency(ut.CurrencyStandard, "USD", 1234.50)
	fmt.Println(formatted)
	// OUTPUT: $1,234.50

	formatted = trans.FmtCurrency(ut.CurrencyStandard, "USD", -1234.50)
	fmt.Println(formatted)
	// OUTPUT: -$1,234.50

	formatted = trans.FmtCurrency(ut.CurrencyAccounting, "USD", -1234.50)
	fmt.Println(formatted)
	// OUTPUT: ($1,234.50)

	formatted = trans.FmtCurrencyWhole(ut.CurrencyStandard, "USD", -1234.50)
	fmt.Println(formatted)
	// OUTPUT: -$1,234

	formatted = trans.FmtNumber(1234.50)
	fmt.Println(formatted)
	// OUTPUT: 1,234.5

	formatted = trans.FmtNumberWhole(1234.50)
	fmt.Println(formatted)
	// OUTPUT: 1,234
}
