package jgo_CM

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type jgo_CM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'jgo_CM' locale
func New() locales.Translator {
	return &jgo_CM{
		locale:  "jgo_CM",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *jgo_CM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'jgo_CM'
func (t *jgo_CM) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'jgo_CM'
func (t *jgo_CM) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
