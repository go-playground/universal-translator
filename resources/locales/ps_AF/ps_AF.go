package ps_AF

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ps_AF struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ps_AF' locale
func New() locales.Translator {
	return &ps_AF{
		locale:  "ps_AF",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ps_AF) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ps_AF'
func (t *ps_AF) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ps_AF'
func (t *ps_AF) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
