package ru_RU

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ru_RU struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ru_RU' locale
func New() locales.Translator {
	return &ru_RU{
		locale:  "ru_RU",
		plurals: []locales.PluralRule{2, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *ru_RU) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ru_RU'
func (t *ru_RU) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ru_RU'
func (t *ru_RU) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if v == 0 && i%10 == 1 && i%100 != 11 {
		return locales.PluralRuleOne
	} else if v == 0 && i%10 >= 2 && i%10 <= 4 && i%100 < 12 && i%100 > 14 {
		return locales.PluralRuleFew
	} else if (v == 0 && i%10 == 0) || (v == 0 && i%10 >= 5 && i%10 <= 9) || (v == 0 && i%100 >= 11 && i%100 <= 14) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}
