package es_GQ

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type es_GQ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'es_GQ' locale
func New() locales.Translator {
	return &es_GQ{
		locale:  "es_GQ",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *es_GQ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'es_GQ'
func (t *es_GQ) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'es_GQ'
func (t *es_GQ) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
