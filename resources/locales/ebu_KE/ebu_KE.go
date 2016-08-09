package ebu_KE

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ebu_KE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ebu_KE' locale
func New() locales.Translator {
	return &ebu_KE{
		locale:  "ebu_KE",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *ebu_KE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ebu_KE'
func (t *ebu_KE) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ebu_KE'
func (t *ebu_KE) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
