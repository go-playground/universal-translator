package asa_TZ

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type asa_TZ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'asa_TZ' locale
func New() locales.Translator {
	return &asa_TZ{
		locale:  "asa_TZ",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *asa_TZ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'asa_TZ'
func (t *asa_TZ) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'asa_TZ'
func (t *asa_TZ) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
