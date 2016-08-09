package nus

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type nus struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'nus' locale
func New() locales.Translator {
	return &nus{
		locale:  "nus",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *nus) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'nus'
func (t *nus) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'nus'
func (t *nus) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
