package ta_IN

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ta_IN struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ta_IN' locale
func New() locales.Translator {
	return &ta_IN{
		locale:  "ta_IN",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ta_IN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ta_IN'
func (t *ta_IN) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ta_IN'
func (t *ta_IN) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
