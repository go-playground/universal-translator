package ha

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ha struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ha' locale
func New() locales.Translator {
	return &ha{
		locale:  "ha",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ha) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ha'
func (t *ha) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ha'
func (t *ha) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
