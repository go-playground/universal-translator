package kea

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type kea struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'kea' locale
func New() locales.Translator {
	return &kea{
		locale:  "kea",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *kea) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'kea'
func (t *kea) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'kea'
func (t *kea) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
