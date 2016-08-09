package vo

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type vo struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'vo' locale
func New() locales.Translator {
	return &vo{
		locale:  "vo",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *vo) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'vo'
func (t *vo) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'vo'
func (t *vo) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
