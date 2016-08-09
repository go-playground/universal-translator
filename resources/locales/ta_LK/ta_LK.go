package ta_LK

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ta_LK struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ta_LK' locale
func New() locales.Translator {
	return &ta_LK{
		locale:  "ta_LK",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ta_LK) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ta_LK'
func (t *ta_LK) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ta_LK'
func (t *ta_LK) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
