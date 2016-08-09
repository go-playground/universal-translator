package ne_NP

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ne_NP struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ne_NP' locale
func New() locales.Translator {
	return &ne_NP{
		locale:  "ne_NP",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ne_NP) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ne_NP'
func (t *ne_NP) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ne_NP'
func (t *ne_NP) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
