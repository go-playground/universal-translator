package es_DO

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type es_DO struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'es_DO' locale
func New() locales.Translator {
	return &es_DO{
		locale:  "es_DO",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *es_DO) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'es_DO'
func (t *es_DO) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'es_DO'
func (t *es_DO) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
