package fr_CA

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type fr_CA struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fr_CA' locale
func New() locales.Translator {
	return &fr_CA{
		locale:  "fr_CA",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fr_CA) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fr_CA'
func (t *fr_CA) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'fr_CA'
func (t *fr_CA) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 0 || i == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
