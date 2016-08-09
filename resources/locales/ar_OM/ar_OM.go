package ar_OM

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ar_OM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ar_OM' locale
func New() locales.Translator {
	return &ar_OM{
		locale:  "ar_OM",
		plurals: []locales.PluralRule{1, 2, 3, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *ar_OM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ar_OM'
func (t *ar_OM) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ar_OM'
func (t *ar_OM) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

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
