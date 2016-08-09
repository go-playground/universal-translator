package naq

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type naq struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'naq' locale
func New() locales.Translator {
	return &naq{
		locale:  "naq",
		plurals: []locales.PluralRule{2, 3, 6},
	}
}

// Locale returns the current translators string locale
func (t *naq) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'naq'
func (t *naq) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'naq'
func (t *naq) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	}

	return locales.PluralRuleOther
}
