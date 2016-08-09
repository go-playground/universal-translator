package az_Latn_AZ

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type az_Latn_AZ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'az_Latn_AZ' locale
func New() locales.Translator {
	return &az_Latn_AZ{
		locale:  "az_Latn_AZ",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *az_Latn_AZ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'az_Latn_AZ'
func (t *az_Latn_AZ) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'az_Latn_AZ'
func (t *az_Latn_AZ) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
