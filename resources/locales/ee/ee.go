package ee

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ee struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ee' locale
func New() locales.Translator {
	return &ee{
		locale:  "ee",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ee) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ee'
func (t *ee) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ee'
func (t *ee) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
