package ig

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ig struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ig' locale
func New() locales.Translator {
	return &ig{
		locale:  "ig",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *ig) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ig'
func (t *ig) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ig'
func (t *ig) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
