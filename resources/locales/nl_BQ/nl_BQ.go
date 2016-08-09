package nl_BQ

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type nl_BQ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'nl_BQ' locale
func New() locales.Translator {
	return &nl_BQ{
		locale:  "nl_BQ",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *nl_BQ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'nl_BQ'
func (t *nl_BQ) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'nl_BQ'
func (t *nl_BQ) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
