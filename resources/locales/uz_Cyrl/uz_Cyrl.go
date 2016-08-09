package uz_Cyrl

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type uz_Cyrl struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'uz_Cyrl' locale
func New() locales.Translator {
	return &uz_Cyrl{
		locale:  "uz_Cyrl",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *uz_Cyrl) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'uz_Cyrl'
func (t *uz_Cyrl) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'uz_Cyrl'
func (t *uz_Cyrl) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
