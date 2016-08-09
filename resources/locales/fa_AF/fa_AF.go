package fa_AF

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type fa_AF struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fa_AF' locale
func New() locales.Translator {
	return &fa_AF{
		locale:  "fa_AF",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fa_AF) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fa_AF'
func (t *fa_AF) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'fa_AF'
func (t *fa_AF) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
