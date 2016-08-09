package ksf_CM

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ksf_CM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ksf_CM' locale
func New() locales.Translator {
	return &ksf_CM{
		locale:  "ksf_CM",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *ksf_CM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ksf_CM'
func (t *ksf_CM) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ksf_CM'
func (t *ksf_CM) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
