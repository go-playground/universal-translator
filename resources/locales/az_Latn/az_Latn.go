package az_Latn

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type az_Latn struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'az_Latn' locale
func New() locales.Translator {
	return &az_Latn{
		locale:  "az_Latn",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *az_Latn) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'az_Latn'
func (t *az_Latn) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'az_Latn'
func (t *az_Latn) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
