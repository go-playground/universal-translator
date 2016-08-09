package wae

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type wae struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'wae' locale
func New() locales.Translator {
	return &wae{
		locale:  "wae",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *wae) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'wae'
func (t *wae) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'wae'
func (t *wae) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
