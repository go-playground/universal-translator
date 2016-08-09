package sq_AL

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type sq_AL struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sq_AL' locale
func New() locales.Translator {
	return &sq_AL{
		locale:  "sq_AL",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *sq_AL) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sq_AL'
func (t *sq_AL) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'sq_AL'
func (t *sq_AL) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
