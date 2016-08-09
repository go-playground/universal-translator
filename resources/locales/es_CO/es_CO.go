package es_CO

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type es_CO struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'es_CO' locale
func New() locales.Translator {
	return &es_CO{
		locale:  "es_CO",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *es_CO) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'es_CO'
func (t *es_CO) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'es_CO'
func (t *es_CO) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
