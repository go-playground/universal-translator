package bg

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type bg struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'bg' locale
func New() locales.Translator {
	return &bg{
		locale:  "bg",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *bg) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'bg'
func (t *bg) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'bg'
func (t *bg) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
