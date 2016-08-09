package ro

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ro struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ro' locale
func New() locales.Translator {
	return &ro{
		locale:  "ro",
		plurals: []locales.PluralRule{2, 4, 6},
	}
}

// Locale returns the current translators string locale
func (t *ro) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ro'
func (t *ro) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ro'
func (t *ro) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	} else if (v != 0) || (n == 0) || (n != 1 && n%100 >= 1 && n%100 <= 19) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}
