package saq

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type saq struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'saq' locale
func New() locales.Translator {
	return &saq{
		locale:  "saq",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *saq) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'saq'
func (t *saq) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'saq'
func (t *saq) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
