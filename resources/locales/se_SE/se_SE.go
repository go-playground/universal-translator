package se_SE

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type se_SE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'se_SE' locale
func New() locales.Translator {
	return &se_SE{
		locale:  "se_SE",
		plurals: []locales.PluralRule{2, 3, 6},
	}
}

// Locale returns the current translators string locale
func (t *se_SE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'se_SE'
func (t *se_SE) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'se_SE'
func (t *se_SE) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	}

	return locales.PluralRuleOther
}
