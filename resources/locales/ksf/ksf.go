package ksf

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ksf struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ksf' locale
func New() locales.Translator {
	return &ksf{
		locale:  "ksf",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *ksf) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ksf'
func (t *ksf) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ksf'
func (t *ksf) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
