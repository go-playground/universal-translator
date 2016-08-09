package it_IT

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type it_IT struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'it_IT' locale
func New() locales.Translator {
	return &it_IT{
		locale:  "it_IT",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *it_IT) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'it_IT'
func (t *it_IT) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'it_IT'
func (t *it_IT) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
