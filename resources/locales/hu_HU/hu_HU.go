package hu_HU

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type hu_HU struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'hu_HU' locale
func New() locales.Translator {
	return &hu_HU{
		locale:  "hu_HU",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *hu_HU) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'hu_HU'
func (t *hu_HU) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'hu_HU'
func (t *hu_HU) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
