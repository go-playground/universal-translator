package km_KH

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type km_KH struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'km_KH' locale
func New() locales.Translator {
	return &km_KH{
		locale:  "km_KH",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *km_KH) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'km_KH'
func (t *km_KH) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'km_KH'
func (t *km_KH) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
