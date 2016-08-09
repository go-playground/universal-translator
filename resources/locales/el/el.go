package el

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type el struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'el' locale
func New() locales.Translator {
	return &el{
		locale:  "el",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *el) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'el'
func (t *el) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'el'
func (t *el) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
