package ky

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ky struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ky' locale
func New() locales.Translator {
	return &ky{
		locale:  "ky",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ky) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ky'
func (t *ky) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ky'
func (t *ky) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
