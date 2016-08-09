package ky_KG

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ky_KG struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ky_KG' locale
func New() locales.Translator {
	return &ky_KG{
		locale:  "ky_KG",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ky_KG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ky_KG'
func (t *ky_KG) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ky_KG'
func (t *ky_KG) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
