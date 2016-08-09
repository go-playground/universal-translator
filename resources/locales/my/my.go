package my

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type my struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'my' locale
func New() locales.Translator {
	return &my{
		locale:  "my",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *my) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'my'
func (t *my) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'my'
func (t *my) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
