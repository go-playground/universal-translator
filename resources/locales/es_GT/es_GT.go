package es_GT

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type es_GT struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'es_GT' locale
func New() locales.Translator {
	return &es_GT{
		locale:  "es_GT",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *es_GT) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'es_GT'
func (t *es_GT) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'es_GT'
func (t *es_GT) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
