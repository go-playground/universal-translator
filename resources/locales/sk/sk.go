package sk

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type sk struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sk' locale
func New() locales.Translator {
	return &sk{
		locale:  "sk",
		plurals: []locales.PluralRule{2, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *sk) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sk'
func (t *sk) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'sk'
func (t *sk) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	} else if i >= 2 && i <= 4 && v == 0 {
		return locales.PluralRuleFew
	} else if v != 0 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}
