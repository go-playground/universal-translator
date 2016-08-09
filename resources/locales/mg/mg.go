package mg

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type mg struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mg' locale
func New() locales.Translator {
	return &mg{
		locale:  "mg",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *mg) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mg'
func (t *mg) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'mg'
func (t *mg) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
