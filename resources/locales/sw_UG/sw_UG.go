package sw_UG

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type sw_UG struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sw_UG' locale
func New() locales.Translator {
	return &sw_UG{
		locale:  "sw_UG",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *sw_UG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sw_UG'
func (t *sw_UG) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'sw_UG'
func (t *sw_UG) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
