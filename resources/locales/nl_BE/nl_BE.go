package nl_BE

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type nl_BE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'nl_BE' locale
func New() locales.Translator {
	return &nl_BE{
		locale:  "nl_BE",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *nl_BE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'nl_BE'
func (t *nl_BE) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'nl_BE'
func (t *nl_BE) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
