package ka_GE

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ka_GE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ka_GE' locale
func New() locales.Translator {
	return &ka_GE{
		locale:  "ka_GE",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ka_GE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ka_GE'
func (t *ka_GE) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ka_GE'
func (t *ka_GE) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
