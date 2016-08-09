package lb

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type lb struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'lb' locale
func New() locales.Translator {
	return &lb{
		locale:  "lb",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *lb) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'lb'
func (t *lb) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'lb'
func (t *lb) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
