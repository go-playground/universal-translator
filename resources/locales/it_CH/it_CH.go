package it_CH

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type it_CH struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'it_CH' locale
func New() locales.Translator {
	return &it_CH{
		locale:  "it_CH",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *it_CH) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'it_CH'
func (t *it_CH) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'it_CH'
func (t *it_CH) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
