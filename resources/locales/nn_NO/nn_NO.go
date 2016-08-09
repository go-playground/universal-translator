package nn_NO

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type nn_NO struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'nn_NO' locale
func New() locales.Translator {
	return &nn_NO{
		locale:  "nn_NO",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *nn_NO) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'nn_NO'
func (t *nn_NO) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'nn_NO'
func (t *nn_NO) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
