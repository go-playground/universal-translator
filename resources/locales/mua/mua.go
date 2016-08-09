package mua

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type mua struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mua' locale
func New() locales.Translator {
	return &mua{
		locale:  "mua",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *mua) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mua'
func (t *mua) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'mua'
func (t *mua) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
