package kok_IN

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type kok_IN struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'kok_IN' locale
func New() locales.Translator {
	return &kok_IN{
		locale:  "kok_IN",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *kok_IN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'kok_IN'
func (t *kok_IN) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'kok_IN'
func (t *kok_IN) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
