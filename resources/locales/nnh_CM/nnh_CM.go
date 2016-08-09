package nnh_CM

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type nnh_CM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'nnh_CM' locale
func New() locales.Translator {
	return &nnh_CM{
		locale:  "nnh_CM",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *nnh_CM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'nnh_CM'
func (t *nnh_CM) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'nnh_CM'
func (t *nnh_CM) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
