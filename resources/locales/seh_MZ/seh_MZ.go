package seh_MZ

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type seh_MZ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'seh_MZ' locale
func New() locales.Translator {
	return &seh_MZ{
		locale:  "seh_MZ",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *seh_MZ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'seh_MZ'
func (t *seh_MZ) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'seh_MZ'
func (t *seh_MZ) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
