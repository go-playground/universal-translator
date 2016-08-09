package mfe_MU

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type mfe_MU struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mfe_MU' locale
func New() locales.Translator {
	return &mfe_MU{
		locale:  "mfe_MU",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *mfe_MU) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mfe_MU'
func (t *mfe_MU) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'mfe_MU'
func (t *mfe_MU) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
