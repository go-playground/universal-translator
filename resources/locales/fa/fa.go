package fa

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type fa struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fa' locale
func New() locales.Translator {
	return &fa{
		locale:  "fa",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fa) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fa'
func (t *fa) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'fa'
func (t *fa) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
