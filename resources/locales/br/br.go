package br

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type br struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'br' locale
func New() locales.Translator {
	return &br{
		locale:  "br",
		plurals: []locales.PluralRule{2, 3, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *br) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'br'
func (t *br) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'br'
func (t *br) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n%10 == 1 && (n%100 != 11 && n%100 != 71 && n%100 != 91) {
		return locales.PluralRuleOne
	} else if n%10 == 2 && (n%100 != 12 && n%100 != 72 && n%100 != 92) {
		return locales.PluralRuleTwo
	} else if n%10 >= 3 && n%10 <= 4 && (n%10 == 9) && (n%100 < 10 && n%100 > 19) || (n%100 < 70 && n%100 > 79) || (n%100 < 90 && n%100 > 99) {
		return locales.PluralRuleFew
	} else if n != 0 && n%1000000 == 0 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}
