package teo_UG

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type teo_UG struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'teo_UG' locale
func New() locales.Translator {
	return &teo_UG{
		locale:  "teo_UG",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *teo_UG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'teo_UG'
func (t *teo_UG) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'teo_UG'
func (t *teo_UG) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
