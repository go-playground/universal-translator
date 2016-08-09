package es_PR

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type es_PR struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'es_PR' locale
func New() locales.Translator {
	return &es_PR{
		locale:  "es_PR",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *es_PR) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'es_PR'
func (t *es_PR) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'es_PR'
func (t *es_PR) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
