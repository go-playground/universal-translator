package smn_FI

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type smn_FI struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'smn_FI' locale
func New() locales.Translator {
	return &smn_FI{
		locale:  "smn_FI",
		plurals: []locales.PluralRule{2, 3, 6},
	}
}

// Locale returns the current translators string locale
func (t *smn_FI) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'smn_FI'
func (t *smn_FI) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'smn_FI'
func (t *smn_FI) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	}

	return locales.PluralRuleOther
}
