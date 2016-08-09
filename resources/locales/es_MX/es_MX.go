package es_MX

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type es_MX struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'es_MX' locale
func New() locales.Translator {
	return &es_MX{
		locale:  "es_MX",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *es_MX) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'es_MX'
func (t *es_MX) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'es_MX'
func (t *es_MX) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
