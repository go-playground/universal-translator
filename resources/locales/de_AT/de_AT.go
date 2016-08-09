package de_AT

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type de_AT struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'de_AT' locale
func New() locales.Translator {
	return &de_AT{
		locale:  "de_AT",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *de_AT) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'de_AT'
func (t *de_AT) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'de_AT'
func (t *de_AT) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
