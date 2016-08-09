package es_IC

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type es_IC struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'es_IC' locale
func New() locales.Translator {
	return &es_IC{
		locale:  "es_IC",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *es_IC) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'es_IC'
func (t *es_IC) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'es_IC'
func (t *es_IC) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
