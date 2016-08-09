package ks

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ks struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ks' locale
func New() locales.Translator {
	return &ks{
		locale:  "ks",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ks) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ks'
func (t *ks) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ks'
func (t *ks) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
