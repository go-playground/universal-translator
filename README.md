## universal-translator
<img align="right" src="https://raw.githubusercontent.com/go-playground/universal-translator/master/logo.png">
![Project status](https://img.shields.io/badge/version-0.9-green.svg)
[![Build Status](https://semaphoreci.com/api/v1/joeybloggs/universal-translator/branches/master/badge.svg)](https://semaphoreci.com/joeybloggs/universal-translator)
[![Coverage Status](https://coveralls.io/repos/github/go-playground/universal-translator/badge.svg?branch=master)](https://coveralls.io/github/go-playground/universal-translator?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-playground/universal-translator)](https://goreportcard.com/report/github.com/go-playground/universal-translator)
[![GoDoc](https://godoc.org/github.com/go-playground/universal-translator?status.svg)](https://godoc.org/github.com/go-playground/universal-translator)
![License](https://img.shields.io/dub/l/vibe-d.svg)

Universal Translator is an i18n Translator for Go/Golang using CLDR data + pluralization rules

Why another i18n library?
--------------------------
I noticed that most libraries out there use static files for translations, which I'm not against just there is not option for coding it inline, 
as well as having formats which do not handle all plural rules, or are overcomplicated. There is also very little in the way of helping the user
know about what plural translations are needed for each language, no easy grouping to say, display all translations for a page...

Features
--------
- [x] Rules added from [CLDR](http://cldr.unicode.org/index/downloads) data
- [x] Use fmt.Sprintf() for translation string parsing
- [x] Add Translations in code
- [x] Prints the supported plural rules for a given translators locale using translator.PrintPluralRules()
- [x] Plural Translations
- [x] Date, Time & DateTime formatting
- [x] Number, Whole Number formatting
- [x] Currency both standard & accounting, formatting i.e. -$1,234.50 vs ($1,234.50)
- [x] Handles BC and AD Dates. i.e. January 2, 300 BC
- [ ] Support loading translations from files
- [ ] Exporting translations to file, mainly for getting them professionally translated
- [ ] Code Generation for translation files -> Go code.. i.e. after it has been professionally translated
- [ ] Printing of grouped translations, i.e. all transations for the homepage
- [ ] Tests for all languages, I need help with this one see below

Full Language tests
--------------------
I could sure use your help adding tests for every language, it is a huge undertaking and I just don't have the free time to do it all at the moment;
any help would be **greatly appreciated!!!!** please see [issue](https://github.com/go-playground/universal-translator/issues/1) for details.

Installation
-----------

Use go get 

```go
go get github.com/go-playground/universal-translator
``` 

or to update

```go
go get -u github.com/go-playground/universal-translator
``` 

Usage
-------
```go
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
```

Help With Tests
---------------
To anyone interesting in helping or contributing, I sure could use some help creating tests for each language.
Please see issue [here](https://github.com/go-playground/universal-translator/issues/1) for details.

Thanks to some help, the following languages have tests:

- [x] en - English US
- [x] th - Thai thanks to @prideloki

Special Thanks
--------------
Special thanks to the following libraries that not only inspired, but that I borrowed a bunch of code from to create this.. ultimately there were many changes made and more to come, but without them would have taken forever to just get started.
* [cldr](https://github.com/theplant/cldr) - A golang i18n tool using CLDR data
* [i18n](https://github.com/vube/i18n) - golang package for basic i18n features, including message translation and number formatting

Misc
-------
Library is not at 1.0 yet, but don't forsee any major API changes; will raise to 1.0 once I've used it completely in at least one project without issue.
