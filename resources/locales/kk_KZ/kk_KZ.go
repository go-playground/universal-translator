package kk_KZ

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type kk_KZ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'kk_KZ' locale
func New() locales.Translator {
	return &kk_KZ{
		locale:  "kk_KZ",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *kk_KZ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'kk_KZ'
func (t *kk_KZ) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'kk_KZ'
func (t *kk_KZ) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
