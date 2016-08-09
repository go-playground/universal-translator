package es_HN

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type es_HN struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'es_HN' locale
func New() locales.Translator {
	return &es_HN{
		locale:  "es_HN",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *es_HN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'es_HN'
func (t *es_HN) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'es_HN'
func (t *es_HN) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
