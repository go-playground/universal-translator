package eu_ES

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type eu_ES struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'eu_ES' locale
func New() locales.Translator {
	return &eu_ES{
		locale:  "eu_ES",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *eu_ES) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'eu_ES'
func (t *eu_ES) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'eu_ES'
func (t *eu_ES) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
