package fy_NL

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type fy_NL struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fy_NL' locale
func New() locales.Translator {
	return &fy_NL{
		locale:  "fy_NL",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fy_NL) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fy_NL'
func (t *fy_NL) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'fy_NL'
func (t *fy_NL) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
