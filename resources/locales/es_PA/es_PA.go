package es_PA

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type es_PA struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'es_PA' locale
func New() locales.Translator {
	return &es_PA{
		locale:  "es_PA",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *es_PA) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'es_PA'
func (t *es_PA) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'es_PA'
func (t *es_PA) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
