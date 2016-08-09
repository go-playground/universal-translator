package es_UY

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type es_UY struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'es_UY' locale
func New() locales.Translator {
	return &es_UY{
		locale:  "es_UY",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *es_UY) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'es_UY'
func (t *es_UY) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'es_UY'
func (t *es_UY) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
