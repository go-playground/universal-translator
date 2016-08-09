package fr_TG

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type fr_TG struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fr_TG' locale
func New() locales.Translator {
	return &fr_TG{
		locale:  "fr_TG",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fr_TG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fr_TG'
func (t *fr_TG) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'fr_TG'
func (t *fr_TG) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 0 || i == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
