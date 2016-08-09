package eo_001

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type eo_001 struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'eo_001' locale
func New() locales.Translator {
	return &eo_001{
		locale:  "eo_001",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *eo_001) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'eo_001'
func (t *eo_001) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'eo_001'
func (t *eo_001) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
