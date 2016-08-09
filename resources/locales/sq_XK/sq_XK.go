package sq_XK

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type sq_XK struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sq_XK' locale
func New() locales.Translator {
	return &sq_XK{
		locale:  "sq_XK",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *sq_XK) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sq_XK'
func (t *sq_XK) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'sq_XK'
func (t *sq_XK) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
