package yav

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type yav struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'yav' locale
func New() locales.Translator {
	return &yav{
		locale:  "yav",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *yav) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'yav'
func (t *yav) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'yav'
func (t *yav) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
