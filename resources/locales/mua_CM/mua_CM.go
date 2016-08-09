package mua_CM

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type mua_CM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mua_CM' locale
func New() locales.Translator {
	return &mua_CM{
		locale:  "mua_CM",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *mua_CM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mua_CM'
func (t *mua_CM) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'mua_CM'
func (t *mua_CM) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
