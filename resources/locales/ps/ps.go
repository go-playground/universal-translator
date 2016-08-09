package ps

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ps struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ps' locale
func New() locales.Translator {
	return &ps{
		locale:  "ps",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ps) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ps'
func (t *ps) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ps'
func (t *ps) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
