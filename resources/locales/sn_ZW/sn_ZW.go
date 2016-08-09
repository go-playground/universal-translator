package sn_ZW

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type sn_ZW struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sn_ZW' locale
func New() locales.Translator {
	return &sn_ZW{
		locale:  "sn_ZW",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *sn_ZW) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sn_ZW'
func (t *sn_ZW) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'sn_ZW'
func (t *sn_ZW) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
