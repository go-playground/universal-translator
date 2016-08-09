package kk

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type kk struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'kk' locale
func New() locales.Translator {
	return &kk{
		locale:  "kk",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *kk) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'kk'
func (t *kk) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'kk'
func (t *kk) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
