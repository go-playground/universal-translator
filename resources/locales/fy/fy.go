package fy

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type fy struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fy' locale
func New() locales.Translator {
	return &fy{
		locale:  "fy",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fy) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fy'
func (t *fy) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'fy'
func (t *fy) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
