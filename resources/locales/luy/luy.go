package luy

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type luy struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'luy' locale
func New() locales.Translator {
	return &luy{
		locale:  "luy",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *luy) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'luy'
func (t *luy) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'luy'
func (t *luy) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
