package az_Cyrl

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type az_Cyrl struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'az_Cyrl' locale
func New() locales.Translator {
	return &az_Cyrl{
		locale:  "az_Cyrl",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *az_Cyrl) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'az_Cyrl'
func (t *az_Cyrl) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'az_Cyrl'
func (t *az_Cyrl) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
