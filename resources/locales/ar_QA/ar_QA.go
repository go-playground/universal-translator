package ar_QA

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ar_QA struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ar_QA' locale
func New() locales.Translator {
	return &ar_QA{
		locale:  "ar_QA",
		plurals: []locales.PluralRule{1, 2, 3, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *ar_QA) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ar_QA'
func (t *ar_QA) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ar_QA'
func (t *ar_QA) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 0 {
		return locales.PluralRuleZero
	} else if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	} else if n%100 >= 3 && n%100 <= 10 {
		return locales.PluralRuleFew
	} else if n%100 >= 11 && n%100 <= 99 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}
