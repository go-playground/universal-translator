package fr_DZ

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type fr_DZ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fr_DZ' locale
func New() locales.Translator {
	return &fr_DZ{
		locale:  "fr_DZ",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fr_DZ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fr_DZ'
func (t *fr_DZ) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'fr_DZ'
func (t *fr_DZ) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 0 || i == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
