package uz_Arab

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type uz_Arab struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'uz_Arab' locale
func New() locales.Translator {
	return &uz_Arab{
		locale:  "uz_Arab",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *uz_Arab) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'uz_Arab'
func (t *uz_Arab) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'uz_Arab'
func (t *uz_Arab) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
