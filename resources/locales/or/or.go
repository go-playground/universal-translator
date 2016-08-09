package or

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type or struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'or' locale
func New() locales.Translator {
	return &or{
		locale:  "or",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *or) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'or'
func (t *or) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'or'
func (t *or) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
