package ewo_CM

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ewo_CM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ewo_CM' locale
func New() locales.Translator {
	return &ewo_CM{
		locale:  "ewo_CM",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *ewo_CM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ewo_CM'
func (t *ewo_CM) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ewo_CM'
func (t *ewo_CM) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
