package fr_SY

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type fr_SY struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fr_SY' locale
func New() locales.Translator {
	return &fr_SY{
		locale:  "fr_SY",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fr_SY) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fr_SY'
func (t *fr_SY) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'fr_SY'
func (t *fr_SY) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 0 || i == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
