package dua_CM

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type dua_CM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'dua_CM' locale
func New() locales.Translator {
	return &dua_CM{
		locale:  "dua_CM",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *dua_CM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'dua_CM'
func (t *dua_CM) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'dua_CM'
func (t *dua_CM) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
