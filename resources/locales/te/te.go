package te

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type te struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'te' locale
func New() locales.Translator {
	return &te{
		locale:  "te",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *te) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'te'
func (t *te) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'te'
func (t *te) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
