package ewo

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ewo struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ewo' locale
func New() locales.Translator {
	return &ewo{
		locale:  "ewo",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *ewo) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ewo'
func (t *ewo) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ewo'
func (t *ewo) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
