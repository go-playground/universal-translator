package km

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type km struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'km' locale
func New() locales.Translator {
	return &km{
		locale:  "km",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *km) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'km'
func (t *km) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'km'
func (t *km) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
