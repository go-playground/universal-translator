package yav_CM

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type yav_CM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'yav_CM' locale
func New() locales.Translator {
	return &yav_CM{
		locale:  "yav_CM",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *yav_CM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'yav_CM'
func (t *yav_CM) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'yav_CM'
func (t *yav_CM) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
