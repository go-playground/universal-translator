package lkt_US

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type lkt_US struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'lkt_US' locale
func New() locales.Translator {
	return &lkt_US{
		locale:  "lkt_US",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *lkt_US) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'lkt_US'
func (t *lkt_US) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'lkt_US'
func (t *lkt_US) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
