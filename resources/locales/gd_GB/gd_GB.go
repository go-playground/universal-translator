package gd_GB

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type gd_GB struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'gd_GB' locale
func New() locales.Translator {
	return &gd_GB{
		locale:  "gd_GB",
		plurals: []locales.PluralRule{2, 3, 4, 6},
	}
}

// Locale returns the current translators string locale
func (t *gd_GB) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'gd_GB'
func (t *gd_GB) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'gd_GB'
func (t *gd_GB) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 || n == 11 {
		return locales.PluralRuleOne
	} else if n == 2 || n == 12 {
		return locales.PluralRuleTwo
	} else if (n >= 3 && n <= 10) || (n >= 13 && n <= 19) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}
