package yue

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type yue struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'yue' locale
func New() locales.Translator {
	return &yue{
		locale:  "yue",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *yue) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'yue'
func (t *yue) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'yue'
func (t *yue) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
