package xog_UG

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type xog_UG struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'xog_UG' locale
func New() locales.Translator {
	return &xog_UG{
		locale:  "xog_UG",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *xog_UG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'xog_UG'
func (t *xog_UG) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'xog_UG'
func (t *xog_UG) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
