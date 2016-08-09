package rn_BI

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type rn_BI struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'rn_BI' locale
func New() locales.Translator {
	return &rn_BI{
		locale:  "rn_BI",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *rn_BI) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'rn_BI'
func (t *rn_BI) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'rn_BI'
func (t *rn_BI) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
