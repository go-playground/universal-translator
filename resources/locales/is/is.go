package is

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type is struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'is' locale
func New() locales.Translator {
	return &is{
		locale:  "is",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *is) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'is'
func (t *is) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'is'
func (t *is) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	t := locales.T(n, v)

	if (t == 0 && i%10 == 1 && i%100 != 11) || (t != 0) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
