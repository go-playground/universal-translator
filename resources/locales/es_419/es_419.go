package es_419

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type es_419 struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'es_419' locale
func New() locales.Translator {
	return &es_419{
		locale:  "es_419",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *es_419) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'es_419'
func (t *es_419) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'es_419'
func (t *es_419) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
