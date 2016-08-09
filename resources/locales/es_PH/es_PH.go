package es_PH

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type es_PH struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'es_PH' locale
func New() locales.Translator {
	return &es_PH{
		locale:  "es_PH",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *es_PH) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'es_PH'
func (t *es_PH) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'es_PH'
func (t *es_PH) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
