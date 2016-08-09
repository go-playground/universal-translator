package fi_FI

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type fi_FI struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fi_FI' locale
func New() locales.Translator {
	return &fi_FI{
		locale:  "fi_FI",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fi_FI) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fi_FI'
func (t *fi_FI) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'fi_FI'
func (t *fi_FI) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
