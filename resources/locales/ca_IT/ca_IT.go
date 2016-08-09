package ca_IT

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ca_IT struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ca_IT' locale
func New() locales.Translator {
	return &ca_IT{
		locale:  "ca_IT",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ca_IT) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ca_IT'
func (t *ca_IT) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ca_IT'
func (t *ca_IT) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
