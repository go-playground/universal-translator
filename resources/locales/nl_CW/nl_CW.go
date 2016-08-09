package nl_CW

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type nl_CW struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'nl_CW' locale
func New() locales.Translator {
	return &nl_CW{
		locale:  "nl_CW",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *nl_CW) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'nl_CW'
func (t *nl_CW) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'nl_CW'
func (t *nl_CW) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
