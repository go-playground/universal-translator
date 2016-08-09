package si_LK

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type si_LK struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'si_LK' locale
func New() locales.Translator {
	return &si_LK{
		locale:  "si_LK",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *si_LK) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'si_LK'
func (t *si_LK) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'si_LK'
func (t *si_LK) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)

	if (n == 0 || n == 1) || (i == 0 && f == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
