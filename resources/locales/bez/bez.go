package bez

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type bez struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'bez' locale
func New() locales.Translator {
	return &bez{
		locale:  "bez",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *bez) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'bez'
func (t *bez) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'bez'
func (t *bez) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
