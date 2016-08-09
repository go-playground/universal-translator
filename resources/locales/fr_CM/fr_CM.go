package fr_CM

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type fr_CM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fr_CM' locale
func New() locales.Translator {
	return &fr_CM{
		locale:  "fr_CM",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fr_CM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fr_CM'
func (t *fr_CM) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'fr_CM'
func (t *fr_CM) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 0 || i == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
