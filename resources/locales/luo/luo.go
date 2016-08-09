package luo

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type luo struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'luo' locale
func New() locales.Translator {
	return &luo{
		locale:  "luo",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *luo) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'luo'
func (t *luo) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'luo'
func (t *luo) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
