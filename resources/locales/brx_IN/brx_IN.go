package brx_IN

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type brx_IN struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'brx_IN' locale
func New() locales.Translator {
	return &brx_IN{
		locale:  "brx_IN",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *brx_IN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'brx_IN'
func (t *brx_IN) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'brx_IN'
func (t *brx_IN) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
