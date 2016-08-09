package fr_MU

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type fr_MU struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fr_MU' locale
func New() locales.Translator {
	return &fr_MU{
		locale:  "fr_MU",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fr_MU) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fr_MU'
func (t *fr_MU) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'fr_MU'
func (t *fr_MU) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 0 || i == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
