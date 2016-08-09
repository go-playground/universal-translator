package dje

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type dje struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'dje' locale
func New() locales.Translator {
	return &dje{
		locale:  "dje",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *dje) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'dje'
func (t *dje) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'dje'
func (t *dje) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
