package es_CR

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type es_CR struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'es_CR' locale
func New() locales.Translator {
	return &es_CR{
		locale:  "es_CR",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *es_CR) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'es_CR'
func (t *es_CR) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'es_CR'
func (t *es_CR) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
