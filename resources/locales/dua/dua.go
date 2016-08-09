package dua

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type dua struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'dua' locale
func New() locales.Translator {
	return &dua{
		locale:  "dua",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *dua) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'dua'
func (t *dua) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'dua'
func (t *dua) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
