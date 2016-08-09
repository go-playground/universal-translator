package nn

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type nn struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'nn' locale
func New() locales.Translator {
	return &nn{
		locale:  "nn",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *nn) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'nn'
func (t *nn) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'nn'
func (t *nn) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
