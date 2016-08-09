package fr_MQ

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type fr_MQ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fr_MQ' locale
func New() locales.Translator {
	return &fr_MQ{
		locale:  "fr_MQ",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fr_MQ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fr_MQ'
func (t *fr_MQ) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'fr_MQ'
func (t *fr_MQ) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 0 || i == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
