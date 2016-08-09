package ta_SG

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ta_SG struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ta_SG' locale
func New() locales.Translator {
	return &ta_SG{
		locale:  "ta_SG",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ta_SG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ta_SG'
func (t *ta_SG) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ta_SG'
func (t *ta_SG) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
