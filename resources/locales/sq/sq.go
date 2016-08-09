package sq

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type sq struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sq' locale
func New() locales.Translator {
	return &sq{
		locale:  "sq",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *sq) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sq'
func (t *sq) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'sq'
func (t *sq) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
