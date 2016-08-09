package ne_IN

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ne_IN struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ne_IN' locale
func New() locales.Translator {
	return &ne_IN{
		locale:  "ne_IN",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ne_IN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ne_IN'
func (t *ne_IN) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ne_IN'
func (t *ne_IN) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
