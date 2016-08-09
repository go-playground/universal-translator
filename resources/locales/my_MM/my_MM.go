package my_MM

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type my_MM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'my_MM' locale
func New() locales.Translator {
	return &my_MM{
		locale:  "my_MM",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *my_MM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'my_MM'
func (t *my_MM) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'my_MM'
func (t *my_MM) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
