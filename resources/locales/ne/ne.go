package ne

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ne struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ne' locale
func New() locales.Translator {
	return &ne{
		locale:  "ne",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ne) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ne'
func (t *ne) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ne'
func (t *ne) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
